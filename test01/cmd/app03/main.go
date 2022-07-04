package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var fileSize = make(chan int64)
var sema = make(chan struct{}, 40)
var done = make(chan struct{})

func main() {
	flag.Parse()
	roots := flag.Args()
	// 如果输入路径为null，则默认当前目录
	if len(roots) == 0 {
		roots = []string{"."}
	}
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var waite sync.WaitGroup
	for _, path := range roots {
		waite.Add(1)
		go walkDir(path, &waite, fileSize)
	}
	ticker := time.NewTicker(3 * time.Second)
	var fileCount int64
	var Size int64

	go func() {
		waite.Wait()
		ticker.Stop()
		close(sema)
		close(fileSize)
	}()

loop:
	for {
		select {
		case <-done:
			for range fileSize {
				//	保证通道不被阻塞
			}
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			Size += size
			fileCount++
		case <-ticker.C:
			formatePrint(fileCount, Size)
		}
	}
	formatePrint(fileCount, Size)
}
func formatePrint(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, info := range dirents(dir) {
		if info.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, info.Name())
			walkDir(subdir, n, fileSize)
		} else {
			fileSize <- info.Size()
		}
	}
}
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token:
	case <-done:
		return nil
	}
	defer func() {
		<-sema
	}()
	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("dirents faild. err: %v", err)
		return nil
	}
	return readDir
}
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

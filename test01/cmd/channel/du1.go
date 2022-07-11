package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

func main() {
	flag.Parse()
	roots := flag.Args()

	fileSize := make(chan int64)
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fmt.Printf("%v", roots)
	var group sync.WaitGroup
	for _, filePath := range roots {
		group.Add(1)
		go walkDir(filePath, &group, fileSize)
	}

	go func() {
		group.Wait()
		close(fileSize)
	}()

	var totalSize, fileCount int64
	var tick <-chan time.Time
	tick = time.Tick(5 * time.Second)
loop:
	for {
		select {
		case size, ok := <-fileSize: //必须使用这种形式接受，有可能fileSize已经关闭，但是仍然取值
			if !ok {
				break loop // fileSize closed
			}
			fileCount++
			totalSize += size
		case <-tick:
			formatPrintf(fileCount, totalSize)
		}
	}
	formatPrintf(fileCount, totalSize)

}
func formatPrintf(fileCount, totalSize int64) {
	fmt.Printf("the file count: %d\tthe file size in total: %.1f GB\n", fileCount, float64(totalSize)/1e9)
}
func walkDir(dir string, group *sync.WaitGroup, fileSizes chan<- int64) {
	defer group.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			group.Add(1)
			subDir := filepath.Join(dir, entry.Name()) // 拼接子目录路径
			walkDir(subDir, group, fileSizes)
		} else {
			fileSizes <- entry.Size() //写入文件大小
		}
	}
}
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}                  // acquire  token
	defer func() { <-sema }()           // release token
	readDir, err := ioutil.ReadDir(dir) //读取目录中所有的文件（包括子目录）
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1 dirents err: %v\n", err)
		return nil
	}
	return readDir
}

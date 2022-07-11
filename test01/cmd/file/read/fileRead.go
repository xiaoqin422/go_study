package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 1.打开文件
	file, err := os.Open("c:\\ProJect\\Go\\study\\test01\\cmd\\file\\test.txt")
	if err != nil {
		log.Fatalf("open file err. %v", err)
	}
	defer file.Close()
	// 2.读取文件内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			log.Printf("file read over.")
			break
		}
		if err != nil {
			log.Fatalf("file read err. %v", err)
		}
		strSlice = append(strSlice, tempSlice[:n]...)
	}
	fmt.Printf("the file content:\n %s", string(strSlice))
}

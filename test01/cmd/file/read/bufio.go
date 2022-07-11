package main

import (
	"bufio"
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
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 每次读取的分隔符为 '\n' 每次读取一行
		if err == io.EOF {                  //可能还有读取返回
			fileStr += str
			log.Printf("file read over.")
			break
		}
		if err != nil {
			log.Fatalf("bufio read err. %v", err)
		}
		fileStr += str
	}
	fmt.Printf("the file content:\n%s", fileStr)
}

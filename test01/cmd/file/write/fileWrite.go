package main

import (
	"log"
	"os"
)

func main() {
	// 1.打开文件
	file, err := os.OpenFile("c:\\ProJect\\Go\\study\\test01\\cmd\\file\\test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 066)
	if err != nil {
		log.Fatalf("open file err. %v", err)
	}
	defer file.Close()
	// 2.写入文件
	_, err = file.WriteString("\r\n直接写入的字符串数据")
	if err != nil {
		log.Fatalf("file write err. %v", err)
	}
	_, err = file.Write([]byte("\r\n字节切片直接写入"))
	if err != nil {
		log.Fatalf("file write err. %v", err)
	}
}

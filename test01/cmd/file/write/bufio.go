package main

import (
	"bufio"
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
	writer := bufio.NewWriter(file)
	writer.WriteString("\r\n这是bufIO的直接写入") //将数据写入缓存
	writer.Flush()                         //将缓存中的内容写入文件
}

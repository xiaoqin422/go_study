package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileStr, err := ioutil.ReadFile("c:\\ProJect\\Go\\study\\test01\\cmd\\file\\test.txt")
	if err != nil {
		log.Fatalf("file read err. %v", err)
	}
	fmt.Printf("the file content:\n%s", string(fileStr))
}

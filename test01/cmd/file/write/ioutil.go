package main

import (
	"io/ioutil"
)

func main() {
	ioutil.WriteFile("c:\\ProJect\\Go\\study\\test01\\cmd\\file\\test.txt", []byte("ioutil直接希入吐"), 066)
}

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go handler(dial)
	for true {
		io.Copy(dial, os.Stdin)
	}
}
func handler(conn net.Conn) {
	io.Copy(os.Stdout, conn)
}

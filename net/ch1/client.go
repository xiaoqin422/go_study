package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var CADDRESS string
var logger *log.Logger

func main() {
	flag.StringVar(&CADDRESS, "address", "localhost:8080", "client connection address")
	flag.Parse()
	logger = log.New(os.Stdout, "[client]", log.Llongfile|log.Ldate|log.Lmicroseconds)
	con, err := net.Dial("tcp", CADDRESS)
	if err != nil {
		logger.Fatalf("client connect fail. err: %v", err)
	}
	logger.Printf("client connect success. address: %v", CADDRESS)
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {

		}
	}(con)
	mustCopy(os.Stdout, con)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		logger.Fatalf("client read connection fail. err: %v", err)
	}
}

package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var loggers *log.Logger
var cADDRESS string

func main() {
	loggers = log.New(os.Stdout, "[client]", log.Llongfile|log.Ldate|log.Lmicroseconds)
	flag.StringVar(&cADDRESS, "address", "localhost:8080", "client address")
	flag.Parse()
	conn, err := net.Dial("tcp", cADDRESS)
	if err != nil {
		loggers.Fatal(err)
	}
	loggers.Printf("client connection success. address: %s\n", cADDRESS)
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
func mustCopy(writer io.Writer, reader io.Reader) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		loggers.Fatal(err)
	}
}

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
	done := make(chan struct{})
	go func() {
		_, err2 := io.Copy(os.Stdout, con)
		if err2 != nil {
			logger.Print(err2)
		}
		done <- struct{}{}
	}()
	mustCopy(con, os.Stdin)
	//类型断言，调用*net.TCPConn的方法CloseWrite()只关闭TCP的写连接
	cw := con.(*net.TCPConn) // cw 的值为conn type为tcpConn
	err = cw.CloseWrite()
	//err = con.Close()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println(<-done)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		logger.Fatalf("client read connection fail. err: %v", err)
	}
}

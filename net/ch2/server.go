package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var logger *log.Logger
var ADDRESS string

func main() {
	logger = log.New(os.Stdout, "[server]", log.Llongfile|log.Ldate|log.Lmicroseconds)
	flag.StringVar(&ADDRESS, "address", "localhost:8080", "server address")
	flag.Parse()
	listen, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("server start success. address: %s\n", ADDRESS)
	for {
		accept, err := listen.Accept()
		if err != nil {
			logger.Fatal(err)
		}
		go ServerHandler(accept)
	}

}
func ServerHandler(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(conn)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		echo(conn, scanner.Text(), 1*time.Second)
	}
}
func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}

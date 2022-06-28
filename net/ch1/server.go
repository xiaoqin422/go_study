package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var ADDRESS string

func main() {
	flag.StringVar(&ADDRESS, "address", "localhost:8080", "server address")
	flag.Parse()
	log.SetPrefix("[server]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds)
	listen, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Fatalf("server listen fail. err: %v", err)
	}
	log.Printf("server start success. address: %v", ADDRESS)
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Fatalf("server accept fail. err: %v", err)
		}
		go ServerHandler(accept)
	}
}
func ServerHandler(con net.Conn) {
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {

		}
	}(con)
	for {
		_, err := io.WriteString(con, time.Now().Format("15:04:05")+"\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}

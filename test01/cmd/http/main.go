package main

import (
	"log"
)

func main() {
	server := NewSdkHttpServer("test-httpServer")
	server.Router("/body/once", readBodyOnce)
	server.Router("/header", header)
	server.Router("/url", url)
	server.Router("/form", form)
	server.Router("/query", query)
	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

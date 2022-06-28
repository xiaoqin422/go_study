package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abord := make(chan []byte)
	go func() {
		bytes := make([]byte, 1)
		os.Stdin.Read(bytes)
		abord <- bytes
	}()
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop() //程序结束是 chanel关闭 防止泄露
	for count := 10; count > 0; count-- {
		select {
		case tick := <-ticker.C:
			fmt.Printf("倒计时select. tick: %v\n", tick)
		case <-abord:
			fmt.Printf("abord\n")
		}
	}

}

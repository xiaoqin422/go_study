package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("hello,world"))
		if err != nil {
			return
		} //字符串转切片
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

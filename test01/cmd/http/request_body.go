package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the header : %v", r.Header)
}

// 使用form必须用parseForm
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form : %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", err)
	}
	fmt.Fprintf(w, "after parse form %v\n", r.Form)
}
func query(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Fprintf(w, "the path query is: %v\n", query)
}

// 很多参数获取不到
func url(w http.ResponseWriter, r *http.Request) {
	url, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, "thr URL %v", string(url))
}

// 只能进行单次读取
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ready body failed: %v", err)
		//	需要返回，不然还会执行后续的代码
		return
	}
	// 类型转换，将 []byte 转换为string
	fmt.Fprintf(w, "read the data: %s \n", string(body))
	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来
		fmt.Fprintf(w, "read the data one more time got errot: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data is:", body)
}

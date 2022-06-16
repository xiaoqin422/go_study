package main

import (
	"fmt"
	"net/http"
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}

	// 构建上下文环境，封装对请求的读取 返回
	ctx := &Context{
		W: w,
		R: r,
	}

	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Printf("%v", req)
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.OkJson(resp)
	// 记录日志，写入响应失败。服务器默认处理返回w
	if err != nil {
		fmt.Printf("写入响应失败: %v", err)
		return
	}
}
func signUpHandlerWithContext(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(ctx.W, "%v", err)
		return
	}
	fmt.Printf("%v", req)
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.OkJson(resp)
	// 记录日志，写入响应失败。服务器默认处理返回w
	if err != nil {
		fmt.Printf("写入响应失败: %v", err)
		return
	}
}

func main() {
	server := NewSdkHttpServer("serverContext")
	server.Router("Get", "/signup", signUpHandlerWithContext)
	err := server.Start(":8080")
	if err != nil {
		return
	}
}

package main

import "net/http"

type Server interface {
	/*
		抽象接口，第一个参数为 路由地址  第二个参数为 handler处理函数
	*/
	Router(method string, pattern string, handlerFunc func(ctx *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedOnMap
}

// Router /*路由注册*/
func (s sdkHttpServer) Router(method string, pattern string, handlerFunc func(ctx *Context)) {
	key := s.handler.key(method, pattern)

	s.handler.handlers[key] = handlerFunc
}

func (s sdkHttpServer) Start(address string) error {
	// 路由初始化
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewSdkHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
		// map 需要进行初始化
		handler: &HandlerBasedOnMap{make(map[string]func(ctx *Context))},
	}
}

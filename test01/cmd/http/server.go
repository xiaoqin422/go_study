package main

import "net/http"

type Server interface {
	Router(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s sdkHttpServer) Router(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewSdkHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

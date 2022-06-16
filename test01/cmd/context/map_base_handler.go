package main

import (
	"net/http"
	"strings"
)

type HandlerBasedOnMap struct {
	// key 应该是 method + url
	handlers map[string]func(ctx *Context)
}

func (h HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handlerFunc, ok := h.handlers[key]; ok {
		handlerFunc(NewSignUpContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		_, err := writer.Write([]byte("Not Found"))
		if err != nil {
			return
		}
	}
}

/*构造拦截器拦截key*/
func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return strings.ToLower(method) + "#" + pattern
}

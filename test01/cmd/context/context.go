package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// ReadJson 空接口，处理任意参数/*
func (c *Context) ReadJson(data interface{}) error {
	// 1 读取请求body
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	// 2 序列化到请求表单
	return json.Unmarshal(body, data)
}

func (c *Context) WriteJson(status int, data interface{}) error {
	respJson, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	if err != nil {
		return err
	}
	c.W.WriteHeader(status)
	return nil
}
func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}
func (c Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}
func NewSignUpContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

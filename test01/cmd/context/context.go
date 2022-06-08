package main

import "encoding/json"

type Context struct {
	Name     string
	Password string
}

func (c *Context) readJson(req interface{}) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	c.Name = string(bytes)
}

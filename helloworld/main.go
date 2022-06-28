package main

import (
	"fmt"
)

func main() {
	a := 1
	b := false
	c := 10.0
	d := complex(10.3, 1)
	e := real(d)
	fmt.Println("hello,world!", a, b, c, d, e, d)
}

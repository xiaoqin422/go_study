package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	name string
	age  int
}
type myInterface interface {
	start(a int)
}

func (p Person) start(a int) {
	fmt.Println(a)
}
func reflectFun(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("类型：%v 类型名称：%v 类型种类： %v\n", v, v.Name(), v.Kind())
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind()
	switch kind {
	case reflect.Int:
		fmt.Printf("Int类型的原始值 %v\n", v.Int())
	case reflect.Float32:
		fmt.Printf("Float32类型的原始值 %v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("Float64类型的原始值 %v\n", v.Float())
	case reflect.String:
		fmt.Printf("String类型的原始值 %v\n", v.String())
	default:
		fmt.Printf("还没有判断这种类型\n")
	}
}
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("这是修改后的值")
	}
}
func main() {
	a := 100
	b := 1.2
	c := true
	var e myInt = 32
	var f = Person{
		name: "张三",
		age:  12,
	}
	reflectFun(a)
	reflectFun(b)
	reflectFun(c)
	reflectFun(e)
	reflectFun(f)
	var h = 25
	reflectFun(&h)
	g := make([]int, 1)
	reflectFun(g)
	i := [...]int{1, 2, 3}
	reflectFun(i)
	var Func func(a int)
	Func = func(a int) {
		fmt.Println(a)
	}
	reflectFun(Func)
	var j myInterface = f
	reflectFun(j)
	m := j.(myInterface)
	reflectFun(m)

	reflectValue(100)
	reflectValue(1.2)
	reflectValue(12.222)
	reflectValue("123213")
	var str string = "hello GoLang"
	fmt.Printf("reflect修改数值。修改前：%s \t", str)
	reflectSetValue(&str)
	fmt.Printf("修改后：%s\n", str)
}

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	Score int
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("GetInfo方法被调用。%v", s)
}
func (s *Student) SetInfo(name string, age, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}
func PrintStructMethods(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("%v,%v\n", t.Kind(), t.Elem().Kind())
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Printf("传入的参数不是一个结构体")
		return
	}
	method := t.Method(0)
	fmt.Printf("%#v \n", method)         //reflect.Method{Name:"GetInfo", PkgPath:"", Type:(*reflect.rtype)(0x1400011c120), Func:reflect.Value{typ:(*reflect.rtype)(0x1400011c120), ptr:(unsafe.Pointer)(0x1400011a038), flag:0x13}, Index:0}
	fmt.Printf("方法名称：%s\n", method.Name) //方法名称：GetInfo
	fmt.Printf("方法类型：%s\n", method.Type) //方法类型：func(main.Student) string
	fmt.Printf("方法体：%s\n", method.Func)  //方法体：%!s(func(main.Student) string=0x104fdb970)

	// 方法执行需要使用值变量
	v := reflect.ValueOf(x)

	s := v.MethodByName("GetInfo").Call(nil)
	fmt.Printf("%v\n", s)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("秦笑笑"))
	params = append(params, reflect.ValueOf(12))
	params = append(params, reflect.ValueOf(99))
	setInfo := v.MethodByName("SetInfo")
	setInfo.Call(params)
	s = v.MethodByName("GetInfo").Call(nil)
	fmt.Printf("%v\n", s)
}
func reflectSetStructFieldValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	} else {
		fmt.Printf("传入的参数不是一个指针类型的结构体\n")
		return
	}
	v.FieldByName("Name").Set(reflect.ValueOf("测试修改"))
	v.FieldByName("Age").Set(reflect.ValueOf(21))
	v.FieldByName("Score").Set(reflect.ValueOf(100))
}
func PrintStructField(x interface{}) {
	// 判断参数是不是结构体类型
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Printf("传入的参数不是一个结构体")
		return
	}

	//1、通过类型变量里边的Field可以获取结构体的字段
	field := t.Field(0)
	fmt.Printf("%#v \n", field)                     //reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x105011820), Tag:"json:\"name\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Printf("字段名称：%s\n", field.Name)             //字段名称：Name
	fmt.Printf("字段类型：%s\n", field.Type)             //字段类型：string
	fmt.Printf("字段Tag：%s\n", field.Tag.Get("json")) //字段Tag：name
	//2、通过类型变量里的FieldByName可以获取到结构体的字段
	//3、通过类型变量里边的NumField获取到结构体有几个字段
	numField := t.NumField()
	var object = make(map[string]interface{}, numField)
	// 获取struct中的所有字段名称
	for i := 0; i < numField; i++ {
		f := t.Field(i)
		key := f.Name
		kin := f.Type
		object[key] = kin
	}

	// 获取传入参数的原始数值
	v := reflect.ValueOf(x)
	fmt.Printf("%v\n", v)      //&{小明 12 99}
	fmt.Printf("%v\n", object) //map[Age:int Name:string Score:int]
	for k, _ := range object {
		value := v.FieldByName(k)
		object[k] = fmt.Sprint(value)
	}
	fmt.Printf("%v\n", object) //map[Age:12 Name:小明 Score:99]

}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   12,
		Score: 99,
	}
	PrintStructField(stu1)   //只是读取属性，不涉及修改
	fmt.Printf("%v\n", stu1) //{小明 12 99}
	fmt.Println("------------------------")
	PrintStructMethods(&stu1)         //涉及到调用指针类型接受参数的方法
	reflectSetStructFieldValue(&stu1) //修改属性，必须是指针类型
	fmt.Printf("%v\n", stu1)          //{测试修改 21 100}

}

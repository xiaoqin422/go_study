package main

import "fmt"

type User struct {
	name    string
	age     int
	address Address
}
type Address struct {
	a string
}

func main() {
	s := make([]int, 0)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	test1(s)
	user := User{
		name:    "qxx",
		age:     23,
		address: Address{a: "ttt"},
	}
	m := make(map[int]*User)
	m[1] = &user
	u := m[1]
	u.name = "qqqqq"
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", user)
	testStruct(user)
	fmt.Printf("%v\n", user)
	err := fmt.Errorf("最少需要配置两个问题答案 %v", []int{1, 2})
	fmt.Println(err.Error())
}
func testStruct(user User) {
	user.name = "111"
	user.address.a = "111"
}
func test1(a []int) {
	fmt.Printf("%v", a)
	for i := 10; i < 20; i++ {
		a = append(a, i)
	}
}
func test2(a string) {
	fmt.Printf("%s\n", a)
}

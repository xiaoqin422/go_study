package main

import "fmt"

var a = complex(1, 6)
var str = "中国人"

func ShareSlice() {
	s1 := []int{1, 2, 3, 4} //定义一个数组
	s2 := s1[2:]            //数组切片
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[1] = 1999
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
}

func main() {
	ShareSlice()
	fmt.Println("hello,world")
	c := real(a)
	d := imag(a)
	fmt.Printf("%X\t%f\t%f\t", a, c, d)
	s := `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
	fmt.Println(s)
	fmt.Println("=============")
	for i := 0; i < len(str); i++ {
		fmt.Printf("index: %d value: 0x%x\n,", i, str[i])
	}
	fmt.Println("=============")
	for i, v := range str {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}
	fmt.Println("=============")
	var list []int
	list = append(list, 1, 2, 3, 4, 5)
	//fmt.Printf("%v", list)
	for i, item := range list {
		fmt.Printf("range🏪 index:%d, value:%d", i, item)
	}
	//for i := 0; i < len(list); i++ {
	//	fmt.Printf("index:%d, value:%d", i, list[i])
	//
	//}
	m := map[int]string{}
	m[1] = "1"
	m[2] = "22222"
	fmt.Println(len(m), m[2])

	for i, s2 := range m {
		fmt.Printf("key: %d, value: %s\n", i, s2)
	}
}

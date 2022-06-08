package main

import (
	"fmt"
)

type ToyDuck struct {
	Color string
	Price int
}

/*
	对象的行为
*/
func (t ToyDuck) Swim() {
	fmt.Printf("Swim方法执行.%v\n", t)
}

func main() {
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()

	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3 是 *ToyDuck
	dock3 := new(ToyDuck) //new 分配好内存，并进行初始化
	dock3.Swim()

	// 当你声明这样的时候，Go 就帮你分配好内存
	// 不用担心空指针的问题，因为它不是指针
	var duck4 ToyDuck
	duck4.Swim()

	// 这是一个指针，但是当前为空指针
	//var duck5 *ToyDuck
	// 会报错
	//duck5.Swim()

	// 推荐使用这种，加入新的属性也可以使用
	duck6 := ToyDuck{
		Color: "黄色",
		Price: 100,
	}
	duck6.Swim()
}

package main

import "fmt"

type treeNode struct {
	Value       int
	Left, Right *treeNode
	a           int
	Node
}
type Node struct {
	ValueN int
}

func (node Node) test() {
	fmt.Println("这是Node struct的方法")
}

func (tree treeNode) test() {
	fmt.Println("这是treeNode struct的方法")
}

func main() {
	tree := new(treeNode)
	tree.Value = 10
	tree.Left = &treeNode{
		Value: 1,
	}
	tree.Right = &treeNode{
		Value: 3,
	}
	tree.ValueN = 1 //匿名结构体可这样访问
	tree.test()
	tree.Node.test()
	test := Node.test // 方法表达式
	test(tree.Node)
	//tra(tree)
}
func tra(t *treeNode) {
	if t == nil {
		return
	}
	tra(t.Left)
	fmt.Println(t.Value)
	tra(t.Right)
}

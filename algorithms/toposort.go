package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string][]string) []string {
	var order []string
	// 记录该节点是否被访问
	seen := make(map[string]bool)
	// 定义遍历函数签名  必须先声明变量,否则无法进行递归
	var visitAll func(items []string)
	// 函数赋值
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// 递归遍历节点之后可以访问的节点
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	// 取出所有的前置节点
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	// 对前置节点进行简单排序
	sort.Strings(keys)
	// 执行函数
	visitAll(keys)
	return order
}

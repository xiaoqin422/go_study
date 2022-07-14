package main

import (
	"demo/model"
	"fmt"
)

var (
	a = model.Q{
		ID:    1,
		Title: "第一个问题",
		As:    []model.A{{ID: 1, Title: "是"}, {ID: 2, Title: "否"}},
	}
	b = model.Q{
		ID:    2,
		Title: "第二个问题",
		As:    []model.A{{ID: 3, Title: "是"}, {ID: 4, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 1,
				Origin_A_ID: 1,
			},
			//{
			//	Origin_Q_ID: 4,
			//	Origin_A_ID: 8,
			//},
		},
	}
	c = model.Q{
		ID:    3,
		Title: "第三个问题",
		As:    []model.A{{ID: 5, Title: "是"}, {ID: 6, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 1,
				Origin_A_ID: 2,
			},
			//{
			//	Origin_Q_ID: 4,
			//	Origin_A_ID: 7,
			//},
		},
	}
	d = model.Q{
		ID:    4,
		Title: "第四个问题",
		As:    []model.A{{ID: 7, Title: "是"}, {ID: 8, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 2,
				Origin_A_ID: 3,
			},
			{
				Origin_Q_ID: 5,
				Origin_A_ID: 9,
			},
		},
	}
	e = model.Q{
		ID:    5,
		Title: "第五个问题",
		As:    []model.A{{ID: 9, Title: "是"}, {ID: 10, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 2,
				Origin_A_ID: 4,
			},
			{
				Origin_Q_ID: 3,
				Origin_A_ID: 5,
			},
		},
	}
	f = model.Q{
		ID:    6,
		Title: "孤立区域",
		As:    []model.A{{ID: 11, Title: "是"}, {ID: 12, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 7,
				Origin_A_ID: 13,
			},
		},
	}
	g = model.Q{
		ID:    7,
		Title: "孤立区域",
		As:    []model.A{{ID: 13, Title: "是"}, {ID: 14, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 6,
				Origin_A_ID: 11,
			},
		},
	}
	h = model.Q{
		ID:    8,
		Title: "孤立区域",
		As:    []model.A{{ID: 15, Title: "是"}, {ID: 16, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 9,
				Origin_A_ID: 17,
			},
		},
	}
	i = model.Q{
		ID:    9,
		Title: "孤立区域",
		As:    []model.A{{ID: 17, Title: "是"}, {ID: 18, Title: "否"}},
		ConditionS: []model.Condition{
			{
				Origin_Q_ID: 8,
				Origin_A_ID: 15,
			},
		},
	}
	data = []model.Q{a, b, c, d, e, f, g, h, i}
)

func main() {
	for _, one := range data {
		conditions := one.ConditionS
		for _, condition := range conditions {
			qId := condition.Origin_Q_ID
			aId := condition.Origin_A_ID
			//origin := db[qId]
			origin := data[qId-1]

			origin.As[(aId+1)%2].Link_Q_ID = one.ID
		}
	}
	visited := make(map[int]int) //图节点访问情况
	errString, valid, result := DFS(1, visited, data)
	if valid {
		fmt.Println(errString)
	}
	fmt.Printf("result:%v", result)
}
func main1() {
	//db := map[int]model.Q{
	//	1: a,
	//	2: b,
	//	3: c,
	//	4: d,
	//	5: e,
	//}
	var last []int
	var errLast []int
	result := make(map[int][]int)
	visited := make(map[int]int) //图节点访问情况
	valid := true
	var dfs func(key int)
	// 构建图结构 o(n*m) m为链接问题
	for _, one := range data {
		conditions := one.ConditionS
		for _, condition := range conditions {
			qId := condition.Origin_Q_ID
			aId := condition.Origin_A_ID
			//origin := db[qId]
			origin := data[qId-1]

			origin.As[(aId+1)%2].Link_Q_ID = one.ID
		}
	}
	dfs = func(key int) {
		visited[key] = 1
		as := data[key-1].As
		for _, ans := range as {
			linkQ := ans.Link_Q_ID
			if linkQ == 0 { //没有下一个问题
				result[key] = append(result[key], ans.ID) //最终问题
			} else {
				if visited[linkQ] == 0 { //这个问题还没有被访问
					dfs(linkQ)
					if !valid {
						errLast = append(errLast, key)
						return
					}
				} else if visited[linkQ] == 1 { //on Path 成环
					valid = false
					// 成环！！！填充环的路径
					errLast = append(errLast, linkQ)
					errLast = append(errLast, key)
					return
				}
			}
		}
		visited[key] = 2
		last = append(last, key)
	}
	dfs(1)
	err1 := make([][]int, 0)
	fmt.Printf("%v\n", result)
	if !valid {
		err1 = append(err1, errLast)
		fmt.Printf("errLast: %v\n", err1)

		return
	}
	// 如果没有遍历完，存在孤立区域
	if len(visited) != len(data) {
		for _, q := range data {
			if visited[q.ID] == 0 {
				dfs(q.ID)
				err1 = append(err1, errLast[:len(errLast)-1]) //环记录追加
				if len(visited) != len(data) {                //一个环校验完毕,继续校验
					errLast = make([]int, 0)
					valid = true
				}
			}
		}
		fmt.Printf("errLast: %v\n", err1)
		return
	}
	fmt.Printf("最终问题：%v\n", result)
	fmt.Printf("访问顺序：  %v\n", last)
	//marshal, err := json.Marshal(data)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%v", string(marshal))
}
func DFS(key int, visited map[int]int, data []model.Q) (errLast string, valid bool, result map[int][]int) {
	//var valid bool
	//var errLast string
	//result := make(map[int][]int, 0)
	visited[key] = 1
	as := data[key-1].As
	for _, ans := range as {
		linkQ := ans.Link_Q_ID
		if linkQ == 0 { //没有下一个问题
			result[key] = append(result[key], ans.ID) //最终问题
		} else {
			if visited[linkQ] == 0 { //这个问题还没有被访问
				errLast, valid, result = DFS(linkQ, visited, data)
				if valid {
					errLast = fmt.Sprintf("%s,%d", errLast, key)
					return errLast, valid, result
				}
			} else if visited[linkQ] == 1 { //on Path 成环
				valid = true
				// 成环！！！填充环的路径
				errLast = fmt.Sprintf("%s,%d,%d", errLast, linkQ, key)
				return errLast, valid, result
			}
		}
	}
	visited[key] = 2
	return errLast, valid, result
}

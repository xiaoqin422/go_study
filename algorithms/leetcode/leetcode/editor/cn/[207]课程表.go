package main

import "fmt"

//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
//示如果要学习课程 ai 则 必须 先学习课程 bi 。
//
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//
//
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
//
//
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
//
//
// 提示：
//
//
// 1 <= numCourses <= 10⁵
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites[i] 中的所有课程对 互不相同
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1326 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		graph   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(n int)
	)
	dfs = func(n int) {
		// 该节点正在遍历，且已经被访问
		visited[n] = 1
		for _, t := range graph[n] {
			if visited[t] == 0 {
				dfs(t)
				if !valid { // 有环后续递归跳出
					return
				}
			} else if visited[t] == 1 { // 该节点正在被遍历，有环
				valid = false
				return
			}
		}
		// 该节点已经被访问过
		visited[n] = 2
		result = append(result, n)
	}
	// 构建图
	for _, info := range prerequisites {
		graph[info[1]] = append(graph[info[1]], info[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	fmt.Printf("%v", result)
	return valid
}

//leetcode submit region end(Prohibit modification and deletion)

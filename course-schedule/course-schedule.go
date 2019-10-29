/**
 * 207. Course Schedule
 *
 * https://leetcode.com/problems/course-schedule/
 */

package leetcode

// 用邻解表的结构去存储图，入度为0表示没有依赖，广度优先遍历图，如果遍历完成后还存在入度为0的结点，则不能完成选课
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)        // 邻解表存储图
	inDegree := make([]int, numCourses) // 所有课程的入度，因为课程使用0-n表示，所以下标可以表示该门课程，值表示入度
	// 构建图和入度
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		inDegree[pre[0]]++
	}

	queue := make([]int, 0) // 用来存放入度为0的结点
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	// 循环遍历 queue
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		for _, v := range graph[head] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// 如果 inDegree 中仍然存在出度不为0的结点，则不能完成选课
	for _, v := range inDegree {
		if v != 0 {
			return false
		}
	}

	return true
}

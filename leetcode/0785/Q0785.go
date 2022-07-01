package leetcode

import "container/list"

func isBipartite(graph [][]int) bool {
	set := make([]int, len(graph))

	q := list.New()
	for i := range graph {
		if set[i] == 0 {
			set[i] = 1
			q.PushBack(i)
			if !bfs(graph, set, q) {
				return false
			}
		}
	}
	return true
}

func bfs(graph [][]int, set []int, q *list.List) bool {
	for q.Len() > 0 {
		from := q.Remove(q.Front()).(int)
		for _, to := range graph[from] {
			setT := -set[from]
			if set[to] == 0 {
				set[to] = setT
				q.PushBack(to)
			} else if set[to] != setT {
				return false
			}
		}
	}
	return true
}

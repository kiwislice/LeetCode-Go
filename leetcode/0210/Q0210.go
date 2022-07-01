package leetcode

import "container/list"

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	incount := make([]int, numCourses)

	queue := list.New()
	for i := range prerequisites {
		from := prerequisites[i][1]
		to := prerequisites[i][0]
		g[from] = append(g[from], to)
		incount[to]++
	}

	for i := range incount {
		if incount[i] == 0 {
			queue.PushBack(i)
		}
	}

	rtn := make([]int, 0, numCourses)
	for queue.Len() > 0 {
		i := queue.Remove(queue.Front()).(int)
		rtn = append(rtn, i)
		subM := g[i]
		if len(subM) > 0 {
			for _, j := range subM {
				incount[j]--
				if incount[j] == 0 {
					queue.PushBack(j)
				}
			}
		}
	}
	if len(rtn) != numCourses {
		rtn = nil
	}
	return rtn
}

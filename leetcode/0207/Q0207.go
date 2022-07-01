package leetcode

import "container/list"

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([]map[int]bool, numCourses)
	incount := make([]int, numCourses)

	for i := range prerequisites {
		from := prerequisites[i][1]
		to := prerequisites[i][0]

		incount[to]++
		subM := g[from]
		if subM == nil {
			subM = make(map[int]bool, numCourses/2)
			g[from] = subM
		}
		subM[to] = true
	}

	// return byDfs(g, numCourses)
	return byBfsTopologicalSort(g, numCourses, incount)
}

func byBfsTopologicalSort(g []map[int]bool, numCourses int, incount []int) bool {
	queue := list.New()
	for i := range incount {
		if incount[i] == 0 {
			queue.PushBack(i)
		}
	}

	count := 0
	for queue.Len() > 0 {
		i := queue.Remove(queue.Front()).(int)
		count++
		subM := g[i]
		if len(subM) > 0 {
			for j := range subM {
				incount[j]--
				if incount[j] == 0 {
					queue.PushBack(j)
				}
			}
		}
	}
	return count == numCourses
}

func byDfs(g []map[int]bool, numCourses int) bool {
	checked := make([]bool, numCourses)
	path := make([]bool, numCourses)

	for i := range g {
		if dfsHasCircle(g, i, checked, path) {
			return false
		}
	}
	return true
}

func dfsHasCircle(g []map[int]bool, start int, checked, path []bool) bool {
	var dfsCircleFound func(i int) bool
	dfsCircleFound = func(i int) bool {
		// fmt.Println(path, i)
		if path[i] {
			return true
		}

		path[i] = true
		if !checked[i] && len(g[i]) > 0 {
			for nextK := range g[i] {
				if dfsCircleFound(nextK) {
					return true
				}
			}
		}
		checked[i] = true
		path[i] = false
		return false
	}
	return dfsCircleFound(start)
}

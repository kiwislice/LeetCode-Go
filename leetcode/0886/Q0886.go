package leetcode

import (
	"container/list"
)

func possibleBipartition(n int, dislikes [][]int) bool {
	G := make([][]bool, n+1)
	for i := range G {
		G[i] = make([]bool, n+1)
	}
	for i := range dislikes {
		from, to := dislikes[i][0], dislikes[i][1]
		G[from][to] = true
		G[to][from] = true
	}

	queue := list.New()
	group := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if group[i] != 0 {
			continue
		}

		queue.PushBack(i)
		var curGroup, nextGroup int
		for queue.Len() > 0 {
			p := queue.Remove(queue.Front()).(int)
			curGroup, nextGroup = sign(p), -sign(p)
			p *= curGroup
			group[p] = curGroup
			for to, hasPath := range G[p] {
				if !hasPath {
					continue
				}
				if group[to] == 0 {
					queue.PushBack(to * nextGroup)
				} else if group[to] != -curGroup {
					return false
				}
			}
		}
	}
	return true
}

func sign(i int) int {
	if i >= 0 {
		return 1
	} else {
		return -1
	}
}

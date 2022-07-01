package leetcode

import (
	"container/list"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	g := make([]*list.List, n)
	for i := 0; i < n; i++ {
		g[i] = list.New()
	}
	for i := range edges {
		a := edges[i][0]
		b := edges[i][1]
		g[a].PushBack(b)
		g[b].PushBack(a)
	}

	longestPath := findLongestPath(g)
	halfLen := len(longestPath) / 2
	if len(longestPath)%2 == 0 {
		return longestPath[halfLen-1 : halfLen+1]
	}
	return longestPath[halfLen : halfLen+1]
}

func findLongestPath(g []*list.List) []int {
	n := len(g)
	longestPath := make([]int, 0, n)
	curPath := make([]int, 0, n)
	inpath := make([]bool, n)

	var dfs func(s int)
	dfs = func(s int) {
		curPath = append(curPath, s)
		inpath[s] = true
		for e := g[s].Front(); e != nil; e = e.Next() {
			v := e.Value.(int)
			if !inpath[v] {
				dfs(e.Value.(int))
			}
		}
		if len(longestPath) < len(curPath) {
			longestPath = append(longestPath[:0], curPath...)
		}
		inpath[curPath[len(curPath)-1]] = false
		curPath = curPath[:len(curPath)-1]
	}
	dfs(0)
	dfs(longestPath[len(longestPath)-1])
	return longestPath
}

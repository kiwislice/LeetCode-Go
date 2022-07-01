package leetcode

func allPathsSourceTarget(graph [][]int) [][]int {
	G := make(map[int]map[int]bool)
	for i := range graph {
		G[i] = make(map[int]bool)
		for _, j := range graph[i] {
			G[i][j] = true
		}
	}

	n := len(graph)
	result := make([][]int, 0, n)
	path := make([]int, 0, n)
	result = dfs(n, G, 0, path, result)
	return result
}

func dfs(n int, G map[int]map[int]bool, from int, path []int, result [][]int) [][]int {
	path = append(path, from)
	if from == n-1 {
		result = append(result, copy(path))
	} else {
		for to, hasPath := range G[from] {
			if hasPath {
				result = dfs(n, G, to, path, result)
			}
		}
	}
	return result
}

func copy(ar []int) []int {
	rtn := make([]int, 0, len(ar))
	rtn = append(rtn, ar...)
	return rtn
}

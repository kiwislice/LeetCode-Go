package leetcode

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	nameToI := make(map[string]int)
	index := 0
	for i := range equations {
		_, exist := nameToI[equations[i][0]]
		if !exist {
			nameToI[equations[i][0]] = index
			index++
		}
		_, exist = nameToI[equations[i][1]]
		if !exist {
			nameToI[equations[i][1]] = index
			index++
		}
	}

	G := make([][]float64, index)
	for i := 0; i < index; i++ {
		G[i] = make([]float64, index)
		G[i][i] = 1.0
	}
	for i := range equations {
		i0 := nameToI[equations[i][0]]
		i1 := nameToI[equations[i][1]]
		G[i0][i1] = values[i]
		G[i1][i0] = 1.0 / values[i]
	}

	rtn := make([]float64, len(queries))
	for i := range queries {
		i0, exist0 := nameToI[queries[i][0]]
		i1, exist1 := nameToI[queries[i][1]]
		if !exist0 || !exist1 {
			rtn[i] = -1.0
			continue
		}
		rtn[i] = findValue(G, i0, i1)
	}
	return rtn
}

func findValue(G [][]float64, from, to int) float64 {
	return dfs(G, from, to, make([]bool, len(G)), 1.0)
}

func dfs(G [][]float64, from, to int, visited []bool, product float64) float64 {
	if G[from][to] > 0.0 {
		return product * G[from][to]
	}

	n := len(G)
	visited[from] = true
	rtn := -1.0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		if G[from][i] > 0.0 {
			rtn = dfs(G, i, to, visited, product*G[from][i])
			if rtn > 0.0 {
				return rtn
			}
		}
	}
	return rtn
}

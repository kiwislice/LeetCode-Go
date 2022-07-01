package leetcode

func findRedundantConnection(edges [][]int) []int {
	n := len(edges) + 1
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}

	findRoot := func(i int) int {
		root := i
		for root != uf[root] {
			uf[root] = uf[uf[root]]
			root = uf[root]
		}
		return root
	}

	for i := range edges {
		ri := findRoot(edges[i][0])
		rj := findRoot(edges[i][1])
		if ri == rj {
			return edges[i]
		}
		uf[ri] = rj
	}
	return edges[len(edges)-1]
}

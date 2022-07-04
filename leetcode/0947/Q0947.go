package leetcode

func removeStones(stones [][]int) int {
	n := len(stones)
	xm := make(map[int]int)
	ym := make(map[int]int)
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}

	findRoot := func(i int) int {
		root := uf[i]
		for root != uf[root] {
			uf[root] = uf[uf[root]]
			root = uf[root]
		}
		return root
	}

	union := func(i, j int) {
		ri := findRoot(i)
		rj := findRoot(j)
		uf[ri] = rj
	}

	for i := range stones {
		x, y := stones[i][0], stones[i][1]

		lastX, existX := xm[x]
		if existX {
			union(lastX, i)
		}
		xm[x] = i

		lastY, existY := ym[y]
		if existY {
			union(lastY, i)
		}
		ym[y] = i
	}

	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		ri := findRoot(i)
		set[ri] = true
	}
	return n - len(set)
}

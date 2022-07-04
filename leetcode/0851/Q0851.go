package leetcode

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	G := make([][]int, n)
	for i := range richer {
		from, to := richer[i][1], richer[i][0]
		if G[from] == nil {
			G[from] = make([]int, 0, n)
		}
		G[from] = append(G[from], to)
	}

	completed := make([]bool, n)
	rtn := make([]int, n)
	for i := 0; i < n; i++ {
		rtn[i] = i
	}
	for i := 0; i < n; i++ {
		dfsMin(G, quiet, i, rtn, completed)
	}
	return rtn
}

func dfsMin(G [][]int, quiet []int, from int, miners []int, completed []bool) (miner int) {
	if completed[from] {
		return miners[from]
	}

	miner = from
	for _, to := range G[from] {
		subMiner := dfsMin(G, quiet, to, miners, completed)
		if quiet[subMiner] < quiet[miner] {
			miner = subMiner
		}
	}
	completed[from] = true
	miners[from] = miner
	return
}

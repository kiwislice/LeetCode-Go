package leetcode

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	uf := newUf(m*n + 1)
	outside := m * n

	rcToI := func(r, c int) int {
		return r*n + c
	}
	// iToRC := func(i int) (r, c int) {
	// 	return i / n, i % n
	// }

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			i := rcToI(r, c)
			if board[r][c] == 'X' || r == 0 || r == m-1 || c == 0 || c == n-1 {
				uf.union(i, outside)
			}
			if c > 0 && board[r][c] == board[r][c-1] {
				uf.union(i, rcToI(r, c-1))
			}
			if r > 0 && board[r][c] == board[r-1][c] {
				uf.union(i, rcToI(r-1, c))
			}
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			i := rcToI(r, c)
			if !uf.isUnion(i, outside) {
				board[r][c] = 'X'
			}
		}
	}

}

type uf struct {
	ps []int
}

func newUf(n int) uf {
	o := uf{ps: make([]int, n)}
	for i := range o.ps {
		o.ps[i] = i
	}
	return o
}

func (o *uf) union(i, j int) {
	ri := o.findRoot(i)
	rj := o.findRoot(j)
	o.ps[ri] = o.ps[rj]
}

func (o *uf) isUnion(i, j int) bool {
	ri := o.findRoot(i)
	rj := o.findRoot(j)
	return ri == rj
}

func (o *uf) findRoot(i int) int {
	root := i
	for root != o.ps[root] {
		o.ps[root] = o.ps[o.ps[root]]
		root = o.ps[root]
	}
	return root
}

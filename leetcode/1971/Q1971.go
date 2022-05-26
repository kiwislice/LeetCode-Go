package leetcode

func validPath(n int, edges [][]int, source int, destination int) bool {
	uf := newUf(n)
	for i := range edges {
		uf.union(edges[i][0], edges[i][1])
	}
	// fmt.Println(uf.ps)
	return uf.isUnion(source, destination)
}

type uf struct {
	ps []int
}

func newUf(n int) uf {
	o := uf{ps: make([]int, n)}
	for i := 0; i < len(o.ps); i++ {
		o.ps[i] = i
	}
	return o
}

func (uf *uf) union(i, j int) {
	ir := uf.findRoot(i)
	jr := uf.findRoot(j)
	uf.ps[jr] = ir
}

func (uf *uf) findRoot(i int) int {
	root := i
	for uf.ps[root] != root {
		uf.ps[root] = uf.ps[uf.ps[root]]
		root = uf.ps[root]
	}
	return root
}

func (uf *uf) isUnion(i, j int) bool {
	return uf.findRoot(i) == uf.findRoot(j)
}

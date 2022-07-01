package leetcode

import "container/heap"

func networkDelayTime(times [][]int, n int, k int) int {
	size := n + 1
	G := make(map[int]map[int]int, size)
	for i := range times {
		s, e := times[i][0], times[i][1]
		_, exist := G[s]
		if !exist {
			G[s] = make(map[int]int, size)
		}
		G[s][e] = times[i][2]
	}

	return byDijkstra(G, size, k)
}

type Distance struct {
	W int
	V int
}

type DistanceHeap []Distance

func (h DistanceHeap) Len() int           { return len(h) }
func (h DistanceHeap) Less(i, j int) bool { return h[i].W < h[j].W }
func (h DistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(Distance))
}

func (h *DistanceHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func byDijkstra(G map[int]map[int]int, size int, k int) int {
	minDist := make([]int, size)
	for i := range minDist {
		minDist[i] = -1
	}

	pq := &DistanceHeap{Distance{0, k}}
	heap.Init(pq)

	for pq.Len() > 0 {
		from := pq.Pop().(Distance)
		for to, dis := range G[from.V] {
			toDist := dis + from.W
			if minDist[to] == -1 || toDist < minDist[to] {
				minDist[to] = toDist
				heap.Push(pq, Distance{toDist, to})
			}
		}
	}

	max := 0
	for i := 1; i < size; i++ {
		if i == k {
			continue
		}
		if minDist[i] < 0 {
			return -1
		}
		if max < minDist[i] {
			max = minDist[i]
		}
	}
	return max
}

func byBfs(G map[int]map[int]int, size int, k int) int {
	minDist := make([]int, size)
	for i := range minDist {
		minDist[i] = -1
	}
	bfs(G, k, minDist, 0)

	max := 0
	for i := 1; i < size; i++ {
		if i == k {
			continue
		}
		if minDist[i] < 0 {
			return -1
		}
		if max < minDist[i] {
			max = minDist[i]
		}
	}
	return max
}

func bfs(G map[int]map[int]int, from int, minDist []int, sumDist int) {
	// fmt.Println("G", from, G[from])
	for to, dis := range G[from] {
		// fmt.Println("from", from, "to", to, dis, sumDist)
		toDist := dis + sumDist
		if minDist[to] == -1 || toDist < minDist[to] {
			minDist[to] = toDist
			bfs(G, to, minDist, dis+sumDist)
		}
	}
	// fmt.Println("from", from, "minDist", minDist)
}

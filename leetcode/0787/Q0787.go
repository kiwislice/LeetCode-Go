package leetcode

import "container/list"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	G := make(map[int]map[int]int, n)
	for i := range flights {
		from, to := flights[i][0], flights[i][1]
		if _, exists := G[from]; !exists {
			G[from] = make(map[int]int, n)
		}
		G[from][to] = flights[i][2]
	}

	// path := make(map[int]bool, n)
	// return dfsMinPrice(G, src, dst, k+1, 0, path)

	// return pqMinPrice(n, G, src, dst, k)
	return bfsMinPrice(n, G, src, dst, k)
}

type point struct {
	v, w, step int
}

func bfsMinPrice(n int, G map[int]map[int]int, src, dst, k int) int {
	min := make([]int, n)
	for i := range min {
		min[i] = -1
	}

	listA := list.New()
	listA.PushBack(point{src, 0, 0})

	for listA.Len() > 0 {
		p := listA.Remove(listA.Front()).(point)
		for to, toW := range G[p.v] {
			newW := p.w + toW
			if min[to] < 0 || newW < min[to] {
				min[to] = newW
				if p.step < k {
					listA.PushBack(point{to, newW, p.step + 1})
				}
			}
		}

	}
	return min[dst]
}

type Distance struct {
	v    int
	w    int
	step int
}

type DistanceHeap struct {
	arr  []*Distance
	comp func(a, b *Distance) int
}

func (h *DistanceHeap) Len() int           { return len(h.arr) }
func (h *DistanceHeap) Less(i, j int) bool { return h.comp(h.arr[i], h.arr[j]) < 0 }
func (h *DistanceHeap) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func (h *DistanceHeap) Push(x interface{}) {
	h.arr = append(h.arr, x.(*Distance))
}

func (h *DistanceHeap) Pop() interface{} {
	value := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	return value
}

// func pqMinPrice(n int, G map[int]map[int]int, src, dst, k int) int {
// 	comp := func(a, b *Distance) int {
// 		return a.w - b.w
// 	}

// 	pq := &DistanceHeap{[]*Distance{{src, 0, 0}}, comp}
// 	heap.Init(pq)
// 	minPrice := make([][]int, k)
// 	for i := range minPrice {
// 		minPrice[i] = make([]int, n)
// 		for j := range minPrice[i] {
// 			minPrice[i][j] = -1
// 		}
// 	}

// 	for pq.Len() > 0 {
// 		from := heap.Pop(pq).(*Distance)
// 		for to, toDis := range G[from.v] {
// 			obj := &Distance{to, from.w + toDis, from.step + 1}
// 			if minPrice[obj.step][obj.v] < 0 || obj.w < minPrice[obj.v] {
// 				minPrice[obj.v] = obj.w
// 			}
// 			if obj.step < k {
// 				heap.Push(pq, obj)
// 			}
// 		}
// 	}

// 	return minPrice[dst]
// }

func dfsMinPrice(G map[int]map[int]int, src, dst, availableStep, sumPrice int, path map[int]bool) int {
	if availableStep >= 0 && src == dst {
		return sumPrice
	}
	if availableStep <= 0 {
		return -1
	}

	path[src] = true
	minPrice := -1
	for to, price := range G[src] {
		if path[to] {
			continue
		}

		newPrice := dfsMinPrice(G, to, dst, availableStep-1, sumPrice+price, path)
		if newPrice >= 0 {
			if minPrice == -1 || (minPrice >= 0 && newPrice < minPrice) {
				minPrice = newPrice
			}
		}
	}
	path[src] = false
	return minPrice
}

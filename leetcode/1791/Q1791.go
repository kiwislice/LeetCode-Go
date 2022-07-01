package leetcode

func findCenter(edges [][]int) int {
	n := len(edges) + 1
	count := make([]int, n+1)

	for i := range edges {
		count[edges[i][0]]++
		count[edges[i][1]]++
	}

	for i := range count {
		if count[i] == n-1 {
			return i
		}
	}
	return -1
}

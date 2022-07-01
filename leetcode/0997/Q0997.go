package leetcode

func findJudge(n int, trust [][]int) int {
	count := make([]int, n+1)

	for i := range trust {
		count[trust[i][0]]++
		count[trust[i][1]]--
	}

	for i := 1; i <= n; i++ {
		if count[i] == -n+1 {
			return i
		}
	}
	return -1
}

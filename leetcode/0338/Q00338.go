package leetcode

func countBits(n int) []int {
	rtn := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&1 > 0 {
			rtn[i] = 1
		}
		rtn[i] += rtn[i>>1]
	}
	return rtn
}

package leetcode

func diagonalSum(mat [][]int) int {
	sum := 0
	n := len(mat)
	for i := 0; i < n; i++ {
		sum += mat[i][i]
		sum += mat[n-i-1][i]
	}
	if n&1 > 0 {
		sum -= mat[n/2][n/2]
	}
	return sum
}

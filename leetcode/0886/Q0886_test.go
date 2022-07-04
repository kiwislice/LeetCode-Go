package leetcode

import (
	"fmt"
	"testing"

	"github.com/kiwislice/LeetCode-Go/leetcode"
	"github.com/stretchr/testify/assert"
)

func Test_Q0886(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = possibleBipartition(4, leetcode.Int2Array("[[1,2],[1,3],[2,4]]"))
	assert.Equal(true, result)

	result = possibleBipartition(5, leetcode.Int2Array("[[1,2],[2,3],[3,4],[4,5],[1,5]]"))
	assert.Equal(false, result)

	fmt.Println("test finished.")
}

var benchmarkResult any

func BenchmarkPossibleBipartition(b *testing.B) {
	var s any
	for i := 0; i < b.N; i++ {
		s = possibleBipartition(4, leetcode.Int2Array("[[1,2],[1,3],[2,4]]"))
	}
	benchmarkResult = s
}

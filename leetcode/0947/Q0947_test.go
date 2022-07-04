package leetcode

import (
	"fmt"
	"testing"

	"github.com/kiwislice/LeetCode-Go/leetcode"
	"github.com/stretchr/testify/assert"
)

func Test_Q0947(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = removeStones(leetcode.Int2Array("[[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]"))
	assert.Equal(5, result)

	result = removeStones(leetcode.Int2Array("[[0,0],[0,2],[1,1],[2,0],[2,2]]"))
	assert.Equal(3, result)

	fmt.Println("test finished.")
}

/** https://www.practical-go-lessons.com/chap-34-benchmarks */
// var benchmarkResult any

// func BenchmarkAnything(b *testing.B) {
// 	var s any
// 	for i := 0; i < b.N; i++ {
// 		s = countBits(2)
// 	}
// 	benchmarkResult = s
// }

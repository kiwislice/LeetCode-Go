package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0310(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	// result := findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
	// assert.Equal([]int{1}, result)

	// result = findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
	// assert.Equal([]int{4, 3}, result)

	result = findMinHeightTrees(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 4}, {4, 5}, {4, 6}, {6, 7}})
	assert.Equal([]int{0}, result)

	fmt.Println("test finished.")
}

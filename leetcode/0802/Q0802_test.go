package leetcode

import (
	"fmt"
	"testing"

	"github.com/kiwislice/LeetCode-Go/leetcode"
	"github.com/stretchr/testify/assert"
)

func Test_Q0802(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}})
	assert.Equal([]int{2, 4, 5, 6}, result)

	result = eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}})
	assert.Equal([]int{4}, result)

	result = eventualSafeNodes(leetcode.ReadJsonFile_Int2Array("test2.json"))
	assert.Equal(leetcode.ReadJsonFile_IntArray("test2_result.json"), result)

	result = eventualSafeNodes(leetcode.ReadJsonFile_Int2Array("test1.json"))
	assert.Equal(leetcode.ReadJsonFile_IntArray("test1_result.json"), result)

	fmt.Println("test finished.")
}

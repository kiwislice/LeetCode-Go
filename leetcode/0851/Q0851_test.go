package leetcode

import (
	"fmt"
	"testing"

	"github.com/kiwislice/LeetCode-Go/leetcode"
	"github.com/stretchr/testify/assert"
)

func Test_Q0851(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = loudAndRich(leetcode.Int2Array("[[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]]"), []int{3, 2, 5, 4, 6, 1, 7, 0})
	assert.Equal([]int{5, 5, 2, 5, 4, 5, 6, 7}, result)

	result = loudAndRich(leetcode.Int2Array("[]"), []int{0})
	assert.Equal([]int{0}, result)

	param := leetcode.ReadJsonFile_Map("test1.json")
	result = loudAndRich(leetcode.ToInt2d(param["richer"]), leetcode.ToInt1d(param["quiet"]))
	assert.Equal([]int{0, 1, 0, 0, 4, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40}, result)

	fmt.Println("test finished.")
}

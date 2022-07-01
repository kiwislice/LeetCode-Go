package leetcode

import (
	"fmt"
	"testing"

	"github.com/kiwislice/LeetCode-Go/leetcode"
	"github.com/stretchr/testify/assert"
)

func Test_Q0841(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = canVisitAllRooms(leetcode.Int2Array("[[1],[2],[3],[]]"))
	assert.Equal(true, result)

	result = canVisitAllRooms(leetcode.Int2Array("[[1,3],[3,0,1],[2],[0]]"))
	assert.Equal(false, result)

	fmt.Println("test finished.")
}

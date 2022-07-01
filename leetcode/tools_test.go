package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IntArray(t *testing.T) {
	assert := assert.New(t)
	var result []int

	result = IntArray("[2,4,5,6]")
	assert.Equal([]int{2, 4, 5, 6}, result)

	result = IntArray("[2, 4,5, 6]")
	assert.Equal([]int{2, 4, 5, 6}, result)

	result = IntArray("[ ]")
	assert.Equal([]int{}, result)

	fmt.Println("test complete.")
}

func Test_Int2Array(t *testing.T) {
	assert := assert.New(t)
	var result [][]int

	result = Int2Array("[[2,4,5,6]]")
	assert.Equal([][]int{{2, 4, 5, 6}}, result)

	result = Int2Array("[[2, 4] ,[5, 6]]")
	assert.Equal([][]int{{2, 4}, {5, 6}}, result)

	result = Int2Array("[ [] ]")
	assert.Equal([][]int{{}}, result)

	result = Int2Array("[  ]")
	assert.Equal([][]int{}, result)

	fmt.Println("test complete.")
}

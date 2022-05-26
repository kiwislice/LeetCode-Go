package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1572(t *testing.T) {
	assert := assert.New(t)

	result := diagonalSum([][]int{{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1}})
	assert.Equal(8, result)

	result = diagonalSum([][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}})
	assert.Equal(25, result)
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q2032(t *testing.T) {
	assert := assert.New(t)

	result := twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3})
	assert.Equal([]int{2, 3}, result)

	result = twoOutOfThree([]int{3, 1}, []int{2, 3}, []int{1, 2})
	assert.Equal([]int{1, 2, 3}, result)

	result = twoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5})
	assert.Equal([]int{}, result)
}

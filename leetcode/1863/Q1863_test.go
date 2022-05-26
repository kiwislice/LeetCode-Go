package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1863(t *testing.T) {
	assert := assert.New(t)

	result := subsetXORSum([]int{5, 1, 6})
	assert.Equal(28, result)

	result = subsetXORSum([]int{3, 4, 5, 6, 7, 8})
	assert.Equal(480, result)
}

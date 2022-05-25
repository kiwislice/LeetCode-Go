package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1636(t *testing.T) {
	assert := assert.New(t)

	result := frequencySort([]int{1, 1, 2, 2, 2, 3})
	assert.Equal([]int{3, 1, 1, 2, 2, 2}, result)
}

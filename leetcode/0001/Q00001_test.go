package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q00001(t *testing.T) {
	assert := assert.New(t)

	result := twoSum([]int{3, 2, 4}, 6)
	assert.Equal([]int{1, 2}, result)
}

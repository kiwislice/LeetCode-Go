package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1791(t *testing.T) {
	assert := assert.New(t)

	result := findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}})
	assert.Equal(1, result)
}

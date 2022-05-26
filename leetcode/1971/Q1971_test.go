package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1971(t *testing.T) {
	assert := assert.New(t)

	result := validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2)
	assert.Equal(true, result)

	result = validPath(10, [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}, 5, 9)
	assert.Equal(true, result)

	result = validPath(10, [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 7, 5)
	assert.Equal(true, result)
}

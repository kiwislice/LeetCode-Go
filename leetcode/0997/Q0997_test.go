package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0997(t *testing.T) {
	assert := assert.New(t)

	result := findJudge(2, [][]int{{1, 2}})
	assert.Equal(2, result)

	result = findJudge(3, [][]int{{1, 3}, {2, 3}})
	assert.Equal(3, result)

	result = findJudge(1, [][]int{})
	assert.Equal(1, result)
}

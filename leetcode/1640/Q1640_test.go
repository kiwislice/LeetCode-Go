package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1640(t *testing.T) {
	assert := assert.New(t)

	result := canFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}})
	assert.Equal(true, result)
}

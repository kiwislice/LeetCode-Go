package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1700(t *testing.T) {
	assert := assert.New(t)

	result := countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1})
	assert.Equal(0, result)

	result = countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1})
	assert.Equal(3, result)
}

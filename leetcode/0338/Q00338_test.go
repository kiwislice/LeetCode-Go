package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q00338(t *testing.T) {
	assert := assert.New(t)

	result := countBits(2)
	assert.Equal([]int{0, 1, 1}, result)
}

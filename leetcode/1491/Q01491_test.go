package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q01491(t *testing.T) {
	assert := assert.New(t)

	result := average([]int{4000, 3000, 1000, 2000})
	assert.Equal(2500.00000, result)
}

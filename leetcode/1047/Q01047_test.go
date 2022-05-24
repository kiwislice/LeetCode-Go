package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q01047(t *testing.T) {
	assert := assert.New(t)

	result := removeDuplicates("abbaca")
	assert.Equal("ca", result)
}

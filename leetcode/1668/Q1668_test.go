package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1668(t *testing.T) {
	assert := assert.New(t)

	result := maxRepeating("ababc", "ab")
	assert.Equal(2, result)

	result = maxRepeating("ababc", "ac")
	assert.Equal(0, result)

	result = maxRepeating("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba")
	assert.Equal(5, result)

	result = maxRepeating("aa", "ac")
	assert.Equal(0, result)

}

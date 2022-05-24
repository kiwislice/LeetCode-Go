package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q01945(t *testing.T) {
	assert := assert.New(t)

	result := getLucky("iiii", 1)
	assert.Equal(36, result)
}

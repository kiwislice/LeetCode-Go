package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q2194(t *testing.T) {
	assert := assert.New(t)

	result := cellsInRange("K1:L2")
	assert.Equal([]string{"K1", "K2", "L1", "L2"}, result)
}

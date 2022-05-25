package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q1812(t *testing.T) {
	assert := assert.New(t)

	result := squareIsWhite("h3")
	assert.Equal(true, result)
}

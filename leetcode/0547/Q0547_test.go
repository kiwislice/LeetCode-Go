package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0547(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
	assert.Equal(2, result)

	fmt.Println("test finished.")
}

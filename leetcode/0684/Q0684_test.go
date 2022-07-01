package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0684(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})
	assert.Equal([]int{1, 4}, result)

	fmt.Println("test finished.")
}

package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0210(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)

	result := findOrder(2, [][]int{{1, 0}})
	assert.Equal([]int{0, 1}, result)

	result = findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	assert.Equal([]int{0, 1, 2, 3}, result)

	fmt.Println("test finished.")
}

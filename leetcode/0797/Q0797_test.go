package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0797(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}})
	assert.ElementsMatch([][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}, result)

	result = allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {}, {4}, {}})
	assert.ElementsMatch([][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 4}}, result)

	result = allPathsSourceTarget([][]int{{2}, {}, {1}})
	assert.ElementsMatch([][]int{{0, 2}}, result)

	fmt.Println("test finished.")
}

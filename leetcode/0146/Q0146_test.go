package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0146(t *testing.T) {
	assert := assert.New(t)

	cmd := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	arg := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	result := make([]int, 1, len(cmd))

	lru := Constructor(arg[0][0])
	for i := 1; i < len(arg); i++ {
		switch cmd[i] {
		case "put":
			lru.Put(arg[i][0], arg[i][1])
			result = append(result, 0)
		case "get":
			i := lru.Get(arg[i][0])
			result = append(result, i)
		default:
		}
	}
	// result := countBits(2)
	assert.Equal([]int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4}, result)
}

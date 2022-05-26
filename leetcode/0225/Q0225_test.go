package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var null any

func Test_Q0225(t *testing.T) {
	assert := assert.New(t)

	cmd := []string{"MyStack", "push", "push", "top", "pop", "empty"}
	arg := [][]int{{}, {1}, {2}, {}, {}, {}}
	result := make([]any, 1, len(cmd))

	lru := Constructor()
	for i := 1; i < len(arg); i++ {
		switch cmd[i] {
		case "push":
			lru.Push(arg[i][0])
			result = append(result, null)
		case "top":
			i := lru.Top()
			result = append(result, i)
		case "pop":
			i := lru.Pop()
			result = append(result, i)
		case "empty":
			i := lru.Empty()
			result = append(result, i)
		default:
		}
	}
	// result := countBits(2)
	assert.Equal([]any{null, null, null, 2, 2, false}, result)
}

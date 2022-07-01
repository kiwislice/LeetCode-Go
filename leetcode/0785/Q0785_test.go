package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0785(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}})
	assert.Equal(false, result)

	result = isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}})
	assert.Equal(true, result)

	result = isBipartite([][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}})
	assert.Equal(false, result)

	result = isBipartite([][]int{{2, 4}, {2, 3, 4}, {0, 1}, {1}, {0, 1}, {7}, {9}, {5}, {}, {6}, {12, 14}, {}, {10}, {}, {10}, {19}, {18}, {}, {16}, {15}, {23}, {23}, {}, {20, 21}, {}, {}, {27}, {26}, {}, {}, {34}, {33, 34}, {}, {31}, {30, 31}, {38, 39}, {37, 38, 39}, {36}, {35, 36}, {35, 36}, {43}, {}, {}, {40}, {}, {49}, {47, 48, 49}, {46, 48, 49}, {46, 47, 49}, {45, 46, 47, 48}})
	assert.Equal(false, result)

	fmt.Println("test finished.")
}

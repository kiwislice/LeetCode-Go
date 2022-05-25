package leetcode

import (
	"testing"

	"github.com/kiwislice/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

const null = ""

func Test_Q0543(t *testing.T) {
	assert := assert.New(t)

	root := structures.CreateTreeNode(1, 2, 3, 4, 5)
	result := diameterOfBinaryTree(root)
	assert.Equal(3, result)

	root = structures.CreateTreeNode(1, 2)
	result = diameterOfBinaryTree(root)
	assert.Equal(1, result)
}

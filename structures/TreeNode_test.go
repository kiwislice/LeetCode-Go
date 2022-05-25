package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const null = ""

func Test_TreeNode(t *testing.T) {
	assert := assert.New(t)

	result := CreateTreeNode(1, null, 2, 3, 4, 5, null, null, null)
	assert.Equal(CreateTreeNode(1, null, 2, 3, 4, 5), result)
}

package leetcode

import (
	"github.com/kiwislice/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	_, _, m := maxDepth(root)
	return m
}

func maxDepth(root *TreeNode) (leftDepth, rightDepth, maxSumDepth int) {
	if root == nil {
		return -1, -1, 0
	}

	ll, lr, lm := maxDepth(root.Left)
	rl, rr, rm := maxDepth(root.Right)

	leftDepth = max(ll, lr) + 1
	rightDepth = max(rl, rr) + 1
	maxSumDepth = max(leftDepth+rightDepth, lm, rm)
	return
}

func max(a ...int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if m < a[i] {
			m = a[i]
		}
	}
	return m
}

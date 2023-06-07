package lc102_test

import (
	"testing"

	"bensivo.com/leetcode/lc102"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder_Basic(t *testing.T) {
	tree := &lc102.TreeNode{
		Val: 1,
		Left: &lc102.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &lc102.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	assert.Equal(t, [][]int{
		{1},
		{2, 3},
	}, lc102.LevelOrder(tree))
}

func TestLevelOrder_Sparse(t *testing.T) {
	tree := &lc102.TreeNode{
		Val: 3,
		Left: &lc102.TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &lc102.TreeNode{
			Val: 20,
			Left: &lc102.TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &lc102.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.Equal(t, [][]int{
		{3},
		{9, 20},
		{15, 7},
	}, lc102.LevelOrder(tree))
}

package lc226_test

import (
	"testing"

	"bensivo.com/leetcode/lc226"
	"github.com/stretchr/testify/assert"
)

func TestInvertTree_SingleNode(t *testing.T) {

	tree := &lc226.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	inverted := lc226.InvertTree(tree)
	assert.Equal(t, tree, inverted)
}

func TestInvertTree_Full(t *testing.T) {

	// Original
	//    1
	//  2    3
	// 4 5  6 7

	// Inverted
	//    1
	//  3    2
	// 7 6  5 4
	tree := &lc226.TreeNode{
		Val: 1,
		Left: &lc226.TreeNode{
			Val: 2,
			Left: &lc226.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &lc226.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &lc226.TreeNode{
			Val: 3,
			Left: &lc226.TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &lc226.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, lc226.BFS(tree))

	lc226.InvertTree(tree)

	assert.Equal(t, []int{1, 3, 2, 7, 6, 5, 4}, lc226.BFS(tree))
}

func TestInvertTree_Sparse(t *testing.T) {

	//    1
	//  2    3
	// 4 nil nil 7
	tree := &lc226.TreeNode{
		Val: 1,
		Left: &lc226.TreeNode{
			Val: 2,
			Left: &lc226.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &lc226.TreeNode{
			Val:  3,
			Left: nil,
			Right: &lc226.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.Equal(t, []int{1, 2, 3, 4, 7}, lc226.BFS(tree))
	lc226.InvertTree(tree)
	assert.Equal(t, []int{1, 3, 2, 7, 4}, lc226.BFS(tree))
}

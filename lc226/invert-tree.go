package lc226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NOTE: in LC, it looks like simple recursive solutions are quicker and use less memory.
func InvertTree(root *TreeNode) *TreeNode {
	// Do a BFS, turnign the tree into an array
	q := []*TreeNode{root}

	for len(q) > 0 {
		// pop from queue
		node := q[0]
		q = q[1:]

		// Swap left and right nodes
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp

		// Push to queue, continuing BFS
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return root
}

func BFS(root *TreeNode) []int {
	q := []*TreeNode{root}
	res := []int{}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		res = append(res, node.Val)

		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return res
}

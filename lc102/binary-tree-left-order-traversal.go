package lc102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iterative implementation
//
// Uses a BFS queue to iterate through the nodes. Everytime you start a new level,
// the queue itself is already a left-first walk of that level.
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := [][]int{}

	for len(q) > 0 {
		level := []int{}
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]

			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}

// Recursive implementation
//
// Slower and uses more memory
func LevelOrder_Recursive(root *TreeNode) [][]int {

	res := [][]int{}

	var populate func(node *TreeNode, level int)
	populate = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(res) < level+1 {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)
		populate(node.Left, level+1)
		populate(node.Right, level+1)
	}

	populate(root, 0)
	return res
}

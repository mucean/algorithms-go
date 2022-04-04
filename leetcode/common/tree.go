package common

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NilTNode = math.MinInt

func CreateTreeNode(nodes []int) *TreeNode {
	l := len(nodes)
	if l == 0 {
		return nil
	}
	treeNodes := make([]*TreeNode, l)
	treeNodes[0] = &TreeNode{Val: nodes[0]}

	left := true
	for i := 1; i < l; i++ {
		rootPos := (i - 1) / 2
		if nodes[i] == NilTNode {
			continue
		}
		treeNodes[i] = &TreeNode{Val: nodes[i]}
		if left {
			treeNodes[rootPos].Left = treeNodes[i]
		} else {
			treeNodes[rootPos].Right = treeNodes[i]
		}
		left = !left
	}

	return treeNodes[0]
}

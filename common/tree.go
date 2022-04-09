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

func TreeLevelScan(head *TreeNode) []int {
	nodes := make([]int, 0)
	heap := []*TreeNode{head}

	for l := len(heap); l > 0; l = len(heap) {
		for i := 0; i < l; i++ {
			nodes = append(nodes, heap[i].Val)
			if heap[i].Left != nil {
				heap = append(heap, heap[i].Left)
			}
			if heap[i].Right != nil {
				heap = append(heap, heap[i].Right)
			}
		}
		heap = heap[l:]
	}

	return nodes
}

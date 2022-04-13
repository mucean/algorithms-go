package tree

import "github.com/mucean/algorithms-go/common"

var Mirror = mirrorHeap

func mirrorRecursive(root *common.TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left

	mirrorRecursive(root.Left)
	mirrorRecursive(root.Right)
}

func mirrorHeap(root *common.TreeNode) {
	if root == nil {
		return
	}
	list := []*common.TreeNode{root}

	var each *common.TreeNode
	for {
		l := len(list)
		if l == 0 {
			break
		}

		for i := 0; i < l; i++ {
			each = list[i]
			each.Left, each.Right = each.Right, each.Left
			if each.Left != nil {
				list = append(list, each.Left)
			}
			if each.Right != nil {
				list = append(list, each.Right)
			}
		}
		list = list[l:]
	}
}

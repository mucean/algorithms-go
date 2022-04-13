package tree

import "github.com/mucean/algorithms-go/common"

var LevelOrder = levelOrder

func levelOrder(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	list := []*common.TreeNode{root}
	nodes := make([]int, 0)
	for {
		l := len(list)
		if l == 0 {
			break
		}
		for i := 0; i < l; i++ {
			nodes = append(nodes, list[i].Val)
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		list = list[l:]
	}

	return nodes
}

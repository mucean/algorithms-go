package tree

import (
	. "github.com/mucean/algorithms-go/common"
)

var PathSum = pathSum

func pathSum(root *TreeNode, target int) [][]int {
	all := make([][]int, 0)
	if root == nil {
		return all
	}

	var recursive func(root *TreeNode, path []int, sum int)

	recursive = func(root *TreeNode, path []int, sum int) {
		path = append(path, root.Val)
		sum += root.Val

		check := true
		if root.Left != nil {
			recursive(root.Left, path, sum)
			check = false
		}
		if root.Right != nil {
			recursive(root.Right, path, sum)
			check = false
		}
		if check {
			if sum == target {
				each := make([]int, len(path))
				copy(each, path)
				all = append(all, each)
			}
		}
	}

	recursive(root, make([]int, 0, 100), 0)

	return all
}

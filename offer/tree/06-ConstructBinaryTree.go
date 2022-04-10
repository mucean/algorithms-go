package tree

import "github.com/mucean/algorithms-go/common"

var ConstructBinaryTreeFromPreMid = constructBinaryTreeFromPreMid

func constructBinaryTreeFromPreMid(pre, mid []int) *common.TreeNode {
	l := len(pre)
	if l == 0 || l != len(mid) {
		return nil
	}

	root := &common.TreeNode{Val: pre[0]}
	for i := 0; i < l; i++ {
		if pre[0] == mid[i] {
			root.Left = constructBinaryTreeFromPreMid(pre[1:i+1], mid[:i])
			root.Right = constructBinaryTreeFromPreMid(pre[i+1:], mid[i+1:])
		}
	}

	return root
}

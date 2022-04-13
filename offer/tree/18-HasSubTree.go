package tree

import "github.com/mucean/algorithms-go/common"

var HasSubTree = hasSubTree

func hasSubTree(r1, r2 *common.TreeNode) bool {
	var f func(s1, s2 *common.TreeNode) bool
	f = func(s1, s2 *common.TreeNode) bool {
		if s1 == nil {
			return s2 == nil
		}
		if s2 == nil {
			return true
		}

		if s1.Val == s2.Val && f(s1.Left, s2.Left) && f(s1.Right, s2.Right) {
			return true
		}

		return f(s1.Left, r2) || f(s1.Right, r2)
	}

	return f(r1, r2)
}

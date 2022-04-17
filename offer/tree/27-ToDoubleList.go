package tree

import (
	. "github.com/mucean/algorithms-go/common"
)

var ToDoubleList = toDoubleListRecursive

func toDoubleList(root *TreeNode) *TreeNode {
	var head *TreeNode
	if root == nil {
		return head
	}
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre == nil {
			pre = root
			head = root
		} else {
			pre.Right = root
			root.Left = pre
			pre = root
		}
		root = root.Right
	}

	return head
}

func toDoubleListRecursive(root *TreeNode) *TreeNode {
	var recursive func(root *TreeNode) (head *TreeNode, tail *TreeNode)

	var eachHead, eachEnd *TreeNode
	recursive = func(root *TreeNode) (head *TreeNode, end *TreeNode) {
		if root == nil {
			return nil, nil
		}
		head, end = recursive(root.Left)
		if end == nil {
			head, end = root, root
		} else {
			end.Right = root
			root.Left = end
			end = root
		}
		eachHead, eachEnd = recursive(root.Right)
		if eachHead == nil {
			return
		}
		end.Right = eachHead
		eachHead.Left = end
		end = eachEnd
		return
	}

	root, _ = recursive(root)
	return root
}

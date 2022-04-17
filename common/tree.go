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
		if nodes[i] != NilTNode {
			treeNodes[i] = &TreeNode{Val: nodes[i]}
			if left {
				treeNodes[rootPos].Left = treeNodes[i]
			} else {
				treeNodes[rootPos].Right = treeNodes[i]
			}
		}
		left = !left
	}

	return treeNodes[0]
}

func TreeLevelScan(head *TreeNode) []int {
	nodes := make([]int, 0)
	if head == nil {
		return nodes
	}
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

func DoPostorderOther(root *TreeNode) []int {
	nodes := make([]int, 0)
	if root == nil {
		return nodes
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		l := len(stack) - 1
		for l >= 0 && (stack[l].Right == nil || stack[l].Right == root) {
			root = stack[l]
			nodes = append(nodes, root.Val)
			stack = stack[:l]
			l--
		}
		if l >= 0 {
			root = stack[l].Right
		} else {
			root = nil
		}
	}

	return nodes
}

func DoPostorder(root *TreeNode) []int {
	nodes := make([]int, 0)
	if root == nil {
		return nodes
	}
	stack := make([]*TreeNode, 0)
	for root != nil {
		if root.Left != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}
		if root.Right != nil {
			stack = append(stack, root)
			root = root.Right
			continue
		}
		nodes = append(nodes, root.Val)
		for {
			l := len(stack)
			if l == 0 {
				root = nil
				break
			}
			if stack[l-1].Right != root && stack[l-1].Right != nil {
				root = stack[l-1].Right
				break
			}
			root = stack[l-1]
			stack = stack[:l-1]
			nodes = append(nodes, root.Val)
		}
	}

	return nodes
}

func DoPostorderRecursive(root *TreeNode) []int {
	nodes := make([]int, 0)
	var recursive func(root *TreeNode)
	recursive = func(root *TreeNode) {
		if root == nil {
			return
		}
		recursive(root.Left)
		recursive(root.Right)
		nodes = append(nodes, root.Val)
	}

	recursive(root)

	return nodes
}

func DoInorderRecursive(root *TreeNode) []int {
	var recursive func(root *TreeNode, nodes []int) []int

	recursive = func(root *TreeNode, nodes []int) []int {
		if root == nil {
			return nodes
		}
		nodes = append(recursive(root.Left, nodes), root.Val)
		return recursive(root.Right, nodes)
	}

	return recursive(root, []int{})
}

func DoInorder(root *TreeNode) []int {
	nodes := make([]int, 0)
	if root == nil {
		return nodes
	}

	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		root = stack[len(stack)-1]
		nodes = append(nodes, root.Val)
		root = root.Right
		stack = stack[:len(stack)-1]
	}

	return nodes
}

func DoPreorderRecursive(root *TreeNode) []int {
	var recursive func(root *TreeNode, nodes []int) []int

	recursive = func(root *TreeNode, nodes []int) []int {
		if root == nil {
			return nodes
		}
		nodes = append(nodes, root.Val)
		nodes = recursive(root.Left, nodes)
		return recursive(root.Right, nodes)
	}

	return recursive(root, []int{})
}

func DoPreorder(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)

	nodes := make([]int, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			nodes = append(nodes, root.Val)
			stack = append(stack, root)
			root = root.Left
			continue
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return nodes
}

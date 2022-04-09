package p100

import (
	. "github.com/mucean/algorithms-go/common"
)

var ZigzagLevelOrder = zigzagLevelOrder

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	heap := []*TreeNode{root}
	res := make([][]int, 0)
	reverse := false
	for {
		l := len(heap)
		if l == 0 {
			break
		}
		each := make([]int, 0, l)
		for i := 0; i < l; i++ {
			each = append(each, heap[i].Val)
			if heap[i].Right != nil {
				heap = append(heap, heap[i].Right)
			}
			if heap[i].Left != nil {
				heap = append(heap, heap[i].Left)
			}
		}
		if !reverse {
			for s, e := 0, l-1; s < e; s++ {
				each[s], each[e] = each[e], each[s]
				e--
			}
		}
		reverse = !reverse
		heap = heap[l:]
		res = append(res, each)
	}

	return res
}

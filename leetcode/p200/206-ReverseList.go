package p200

import . "github.com/mucean/algorithms-go/leetcode/common"

var ReverseList = reverseList

func reverseList(head *ListNode) *ListNode {
	var res, tmp *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = res
		res = head
		head = tmp
	}

	return res
}

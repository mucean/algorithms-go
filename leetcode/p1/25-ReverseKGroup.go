package p1

import . "github.com/mucean/algorithms-go/leetcode/common"

var ReverseKGroup = reverseKGroup

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	var res, preEnd, eachEnd, eachHead, tmp *ListNode

	for head != nil {
		i := 0
		eachEnd = head
		for ; i < k && head != nil; i++ {
			tmp = head.Next
			head.Next = eachHead
			eachHead = head
			head = tmp
		}

		if i < k {
			head = eachHead
			eachHead = nil
			for ; i > 0; i-- {
				tmp = head.Next
				head.Next = eachHead
				eachHead = head
				head = tmp
			}
		}

		if res == nil {
			res = eachHead
		} else {
			preEnd.Next = eachHead
		}

		eachHead = nil
		preEnd = eachEnd
	}

	return res
}

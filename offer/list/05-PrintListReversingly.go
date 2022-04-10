package list

import "github.com/mucean/algorithms-go/common"

var PrintListReversing = printListReversingIteratively

func printListReversingIteratively(head *common.ListNode) []int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}

	h, e := 0, len(ans)-1
	for h < e {
		ans[h], ans[e] = ans[e], ans[h]
		h++
		e--
	}

	return ans
}

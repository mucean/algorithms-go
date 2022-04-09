package list_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/common"
	. "github.com/mucean/algorithms-go/offer/list"
)

var _ = Describe("List", func() {
	Context("005-PrintListReversing", func() {
		DescribeTable(
			"PrintListReversing",
			func(head *ListNode, res []int) {
				Expect(PrintListReversing(head)).To(Equal(res))
			},
			Entry("nil head", (*ListNode)(nil), []int{}),
			Entry("one node", CreateListNode([]int{1}), []int{1}),
			Entry("normal", CreateListNode([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}),
		)
	})
})

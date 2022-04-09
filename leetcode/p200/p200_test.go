package p200_test

import (
	. "github.com/mucean/algorithms-go/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/leetcode/p200"
)

var _ = Describe("P200", func() {
	DescribeTable(
		"p206",
		func(nodes []int, res []int) {
			list := ReverseList(CreateListNode(nodes))
			var ns []int
			for list != nil {
				ns = append(ns, list.Val)
				list = list.Next
			}
			Expect(ns).To(Equal(res))
		},
		Entry("nil node", []int(nil), []int(nil)),
		Entry("nodes", []int{1, 2, 3}, []int{3, 2, 1}),
	)

	DescribeTable(
		"p215",
		func(nums []int, k, res int) {
			Expect(FindKthLargest(nums, k)).To(Equal(res))
		},
		Entry("only one", []int{22}, 1, 22),
		Entry("min", []int{22, 1, 5, 8}, 4, 1),
		Entry("max", []int{22, 1, 5, 8}, 1, 22),
		Entry("common1", []int{3, 2, 1, 5, 6, 4}, 2, 5),
		Entry("common2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4),
	)
})

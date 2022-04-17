package p50_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/leetcode/p50"
)

var _ = Describe("P50~P99", func() {
	Context("P88", func() {
		DescribeTable(
			"merge two sorted list",
			func(nums1 []int, m int, nums2 []int, n int, ans []int) {
				Merge(nums1, m, nums2, n)
				Expect(nums1).To(Equal(ans))
			},
			Entry("empty num2", []int{1, 2, 3, 0, 0, 0}, 3, []int{}, 0, []int{1, 2, 3, 0, 0, 0}),
			Entry("normal", []int{1, 5, 8, 0, 0, 0}, 3, []int{2, 4, 9}, 3, []int{1, 2, 4, 5, 8, 9}),
			Entry("normal", []int{1, 5, 8, 0, 0, 0}, 3, []int{0, 4, 9}, 3, []int{0, 1, 4, 5, 8, 9}),
			Entry("normal", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}),
			Entry("normal", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}),
		)
	})

	Context("P88-MinDistance", func() {
		DescribeTable(
			"min operation by convert one word to other word",
			func(word1, word2 string, res int) {
				Expect(MinDistance(word1, word2)).To(Equal(res))
			},
			Entry("3", "horse", "ros", 3),
			Entry("5", "intention", "execution", 5),
		)
	})
})

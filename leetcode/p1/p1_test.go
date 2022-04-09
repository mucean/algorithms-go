package p1_test

import (
	. "github.com/mucean/algorithms-go/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/leetcode/p1"
)

var _ = Describe("P1-50", func() {
	Context("P1", func() {
		DescribeTable(
			"two sum",
			func(nums []int, target int, res []int) {
				Expect(TwoSum(nums, target)).To(Equal(res))
			},
			Entry("nil nums", []int(nil), 1, []int(nil)),
			Entry("target not found", []int{1}, 1, []int(nil)),
			Entry("target not found", []int{1, 2, 3}, 1, []int(nil)),
			Entry("normal", []int{2, 7, 11, 15}, 9, []int{0, 1}),
			Entry("normal", []int{3, 2, 4}, 6, []int{1, 2}),
			Entry("normal", []int{3, 3}, 6, []int{0, 1}),
		)
	})

	Context("P3", func() {
		It("test cases", func() {
			Expect(LengthOfLongestSubString("abcabcbb")).To(Equal(3))
			Expect(LengthOfLongestSubString("bbbbbb")).To(Equal(1))
			Expect(LengthOfLongestSubString("pwwkew")).To(Equal(3))
		})
	})

	Context("P15", func() {
		DescribeTable(
			"three sum",
			func(nums []int, ans [][]int) {
				Expect(ThreeSum(nums)).To(Equal(ans))
			},
			Entry("empty nums", []int(nil), [][]int{}),
			Entry("nums too rare", []int{0}, [][]int{}),
			Entry("normal", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}),
		)
	})

	Context("P25", func() {
		DescribeTable(
			"test cases",
			func(nodes []int, k int, res []int) {
				Expect(func() []int {
					var r []int
					n := ReverseKGroup(CreateListNode(nodes), k)
					for n != nil {
						r = append(r, n.Val)
						n = n.Next
					}
					return r
				}()).To(Equal(res))
			},
			Entry("nil node", []int(nil), 2, []int(nil)),
			Entry("k == 1", []int{1, 2, 3}, 1, []int{1, 2, 3}),
			Entry("node length % k == 0", []int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}),
			Entry("node length % k == 1", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}),
		)
	})
})

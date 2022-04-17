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

	Context("P31-NextPermutation", func() {
		DescribeTable(
			"NextPermutation",
			func(nums, res []int) {
				NextPermutation(nums)
				Expect(nums).To(Equal(res))
			},
			Entry("zero element", []int{}, []int{}),
			Entry("one element", []int{1}, []int{1}),
			Entry("two element", []int{1, 2}, []int{2, 1}),
			Entry("two element", []int{2, 1}, []int{1, 2}),
			Entry("three element", []int{1, 2, 3}, []int{1, 3, 2}),
			Entry("three element", []int{1, 3, 2}, []int{2, 1, 3}),
			Entry("three element", []int{3, 2, 1}, []int{1, 2, 3}),
			Entry("three element", []int{2, 3, 1}, []int{3, 1, 2}),
			Entry("five elements", []int{1, 0, 0, 1, 1}, []int{1, 0, 1, 0, 1}),
			Entry("six elements", []int{1, 0, 0, 4, 3, 2}, []int{1, 0, 2, 0, 3, 4}),
			Entry("two elements", []int{1, 1}, []int{1, 1}),
			Entry("more elements", []int{1, 0, 2, 8, 4, 3, 2, 0}, []int{1, 0, 3, 0, 2, 2, 4, 8}),
		)
	})

	Context("P46-Permute", func() {
		DescribeTable(
			"Permute",
			func(nums []int, res [][]int) {
				Expect(Permute(nums)).To(Equal(res))
			},
			Entry("1, 2, 3", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}),
		)
	})

	Context("P47-PermuteUnique", func() {
		DescribeTable(
			"PermuteUnique",
			func(nums []int, res [][]int) {
				Expect(PermuteUnique(nums)).To(Equal(res))
			},
			Entry("1, 1", []int{1, 1}, [][]int{{1, 1}}),
			Entry("1, 1, 1", []int{1, 1, 1}, [][]int{{1, 1, 1}}),
			Entry("1, 3, 1", []int{1, 3, 1}, [][]int{{1, 1, 3}, {1, 3, 1}, {3, 1, 1}}),
		)
	})
})

package array_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/offer-killer/array"
)

var _ = Describe("Array", func() {
	Context("TwoSum func", func() {
		It("test cases", func() {
			Expect(TwoSum([]int{1, 2, 4, 6, 10}, 8)).To(Equal([]int{1, 3}))
			Expect(TwoSum([]int{1, 2}, 3)).To(Equal([]int{0, 1}))
			Expect(TwoSum([]int{1}, 3)).To(Equal([]int(nil)))
		})
	})

	Context("007-ThreeSum", func() {
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

	Context("008-MinSubArrayLen", func() {
		DescribeTable(
			"MinSubArrayLen",
			func(nums []int, target int, res int) {
				Expect(MinSubArrayLen(nums, target)).To(Equal(res))
			},
			Entry("", []int{5, 1, 4, 3}, 7, 2),
		)
	})

	Context("009-SubarrayProductLessThanK", func() {
		DescribeTable(
			"SubarrayProductLessThanK",
			func(nums []int, target int, res [][]int) {
				Expect(SubarrayProductLessThanK(nums, target)).To(Equal(res))
			},
			Entry("normal", []int{5, 1, 4, 3}, 7, [][]int{{5}, {5, 1}, {1}, {1, 4}, {4}, {3}}),
		)
	})

	Context("010-SubarraySumCount", func() {
		DescribeTable(
			"SubarraySumCount",
			func(nums []int, target int, res int) {
				Expect(SubarraySumCount(nums, target)).To(Equal(res))
			},
			Entry("normal", []int{1, 1, 1}, 2, 2),
		)
	})

	Context("011-FindMaxLength", func() {
		DescribeTable(
			"FindMaxLength",
			func(nums []int, res int) {
				Expect(FindMaxLength(nums)).To(Equal(res))
			},
			Entry("only one element", []int{0}, 0),
			Entry("only two element", []int{0, 0}, 0),
			Entry("only two element", []int{0, 1}, 2),
			Entry("normal", []int{0, 1, 0, 0, 1, 1}, 6),
			Entry("normal", []int{0, 0, 1, 0, 0, 1, 1}, 6),
		)
	})

	Context("012-PivotIndex", func() {
		DescribeTable(
			"PivotIndex",
			func(nums []int, res int) {
				Expect(res).To(Equal(PivotIndex(nums)))
			},
			Entry("only one element", []int{1}, 0),
			Entry("no found", []int{1, 2}, -1),
		)
	})

	Context("013-SubMatrixSum", func() {
		Context("normal", func() {
			var matrix = NewNumMatrix([][]int{{3, 0, 1, 2}, {1, 2, 8, 3}, {1, 3, 4, 5}})
			DescribeTable(
				"SubMatrixSum",
				func(row1, col1, row2, col2 int, res int) {
					Expect(matrix.SubMatrixSum(row1, col1, row2, col2)).To(Equal(res))
				},
				Entry("first", 0, 0, 0, 0, 3),
				Entry("four", 0, 0, 1, 1, 6),
				Entry("four", 1, 1, 2, 2, 17),
			)
		})
	})
})

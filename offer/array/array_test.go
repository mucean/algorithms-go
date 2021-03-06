package array_test

import (
	"sort"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/offer/array"
)

var _ = Describe("Array", func() {
	Context("003-FindTwoOrderedDimensionalArray", func() {
		When("array is nil or zero length", func() {
			It("test", func() {
				Expect(FindTwoOrderedDimensionalArray(nil, 1)).To(Equal(false))
				Expect(FindTwoOrderedDimensionalArray([][]int{}, 1)).To(Equal(false))
			})
		})
		DescribeTable(
			"when array is right",
			func(target int, res bool) {
				Expect(FindTwoOrderedDimensionalArray([][]int{
					{1, 2, 8, 9},
					{2, 4, 9, 12},
					{4, 7, 10, 13},
					{6, 8, 11, 15},
				}, target)).To(Equal(res))
			},
			Entry("found: target == min", 1, true),
			Entry("found: target == max", 15, true),
			Entry("found: target == middle", 4, true),
			Entry("found: target == middle", 7, true),
			Entry("not found: target < min", 0, false),
			Entry("not found: target > max", 19, false),
			Entry("not found: target == middle", 5, false),
		)
	})

	DescribeTable(
		"08-Min",
		func(nums []int, res int) {
			Expect(MinArray(nums)).To(Equal(res))
		},
		Entry("normal", []int{5, 6, 7, 0, 1, 2, 3, 4}, 0),
		Entry("normal", []int{5, 6, 7, 1, 2, 3, 4}, 1),
		Entry("normal", []int{1, 1, 1, 1, 1, 1, 0, 1, 1}, 0),
		Entry("normal", []int{1, 2, 3, 4, 5}, 1),
		Entry("normal", []int{3, 4, 5, 6, 7, 2}, 2),
		Entry("normal", []int{0, 1}, 0),
		Entry("normal", []int{1, 0}, 0),
		Entry("normal", []int{0}, 0),
	)

	Context("20-SpiralOrder", func() {
		DescribeTable(
			"SpiralOrder",
			func(matrix [][]int, res []int) {
				Expect(SpiralOrder(matrix)).To(Equal(res))
			},
			Entry("matrix: nil", [][]int(nil), []int{}),
			Entry("matrix: one row", [][]int{{1, 2, 3}}, []int{1, 2, 3}),
			Entry("matrix: one column", [][]int{{1}, {2}, {3}}, []int{1, 2, 3}),
			Entry("matrix: normal", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}),
			Entry("matrix: normal", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}),
		)
	})

	Context("29-MoreThanHalfNum", func() {
		DescribeTable(
			"MoreThanHalfNum",
			func(nums []int, ans int) {
				Expect(MoreThanHalfNum(nums)).To(Equal(ans))
			},
			Entry("nil nums", []int(nil), 0),
			Entry("the array has only one element", []int{1}, 1),
			Entry("the array has only three elements", []int{2, 2, 1}, 2),
		)
	})

	Context("30-KMinNumbers", func() {
		DescribeTable(
			"KMinNumbers",
			func(nums []int, k int, ans []int) {
				res := KMinNumbers(nums, k)
				sort.Ints(res)
				Expect(res).To(Equal(ans))
			},
			Entry("empty", []int{}, 1, []int{}),
			Entry("K great than the length of numbers", []int{5, 2}, 3, []int{2, 5}),
			Entry("normal", []int{5, 2}, 1, []int{2}),
			Entry("more numbers", []int{2, 5, 2, 1, 8, 20, 1}, 3, []int{1, 1, 2}),
		)
	})

	Context("31-GreatestSumOfSubArray", func() {
		DescribeTable(
			"GreatestSumOfSubArray",
			func(nums []int, ans int) {
				Expect(GreatestSumOfSubArray(nums)).To(Equal(ans))
			},
			Entry("nil array", []int(nil), 0),
			Entry("one element in array", []int{1}, 1),
			Entry("two element in array", []int{-1, 1}, 1),
			Entry("two element in array", []int{-5, -2}, -2),
			Entry("normal", []int{1, -2, 3, 10, -4, 7, 2, -5}, 18),
		)
	})
})

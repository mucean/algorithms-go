package array_test

import (
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
})

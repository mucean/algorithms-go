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
})

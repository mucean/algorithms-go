package integer_test

import (
	"math"

	. "github.com/mucean/algorithms-go/offer-killer/integer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integer test suite", func() {
	Describe("Divide func: only by minus and plus", func() {
		Context("overflow int32 max value", func() {
			It("should return max int32 value", func() {
				Expect(Divide(math.MinInt32, -1)).To(Equal(int32(math.MaxInt32)))
			})
		})
	})

	Context("BinarySum func", func() {
		It("", func() {
			Expect(BinarySum("1", "1")).To(Equal("10"))
			Expect(BinarySum("1", "0")).To(Equal("1"))
			Expect(BinarySum("111", "1")).To(Equal("1000"))
		})
	})

	Context("CountBits func", func() {
		It("test cases", func() {
			Expect(CountBits(0)).To(Equal([]int{0}))
			Expect(CountBits(1)).To(Equal([]int{0, 1}))
			Expect(CountBits(2)).To(Equal([]int{0, 1, 1}))
			Expect(CountBits(3)).To(Equal([]int{0, 1, 1, 2}))
			Expect(CountBits(4)).To(Equal([]int{0, 1, 1, 2, 1}))
			Expect(CountBits(5)).To(Equal([]int{0, 1, 1, 2, 1, 2}))
			Expect(CountBits(7)).To(Equal([]int{0, 1, 1, 2, 1, 2, 2, 3}))
		})
	})

	Context("SingleNumber func", Focus, func() {
		It("test cases", func() {
			Expect(SingleNumber([]int{3, 2, 2, 2})).To(Equal(3))
			Expect(SingleNumber([]int{3, 2, 2, 2, 1, 3, 3})).To(Equal(1))
			Expect(SingleNumber([]int{5})).To(Equal(5))
		})
	})

	Context("MaxProduct func", func() {
		It("test cases", func() {
			Expect(MaxProduct([]string{"abcd", "bcd", "efg", "ixyz"})).To(Equal(16))
			Expect(MaxProduct([]string{"abcd", "bcd"})).To(Equal(0))
			Expect(MaxProduct([]string{"bcd"})).To(Equal(0))
		})
	})
})

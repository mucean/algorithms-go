package num_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/offer/num"
)

var _ = Describe("Num", func() {
	Context("09-1-fibonacii", func() {
		DescribeTable(
			"fibonacii",
			func(n, res int) {
				Expect(Fibonacii(n)).To(Equal(res))
			},
			Entry("0", 0, 0),
			Entry("1", 1, 1),
			Entry("4", 4, 3),
			Entry("4", 8, 21),
		)
	})

	Context("11-MyPow", func() {
		DescribeTable(
			"MyPow",
			func(x float64, exp int, ans float64) {
				Expect(FloatEqual(MyPow(x, exp), ans)).To(Equal(true))
			},
			Entry("zero", float64(0), 1, float64(0)),
			Entry("pow zero", float64(100), 0, float64(1)),
			Entry("2 pow 10", float64(2), 10, float64(1024)),
			Entry("2.1 pow 3", 2.1, 3, 9.26100),
			Entry("2 pow -2", float64(2), -2, 0.2500),
		)
	})
})

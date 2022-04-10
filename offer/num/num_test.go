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
})

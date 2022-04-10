package binary_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/offer/binary"
)

var _ = Describe("Binary", func() {
	Context("10-NumberOf1", func() {
		DescribeTable(
			"NumberOf1",
			func(n, res int) {
				Expect(NumberOf1(n)).To(Equal(res))
			},
			Entry("0", 0, 0),
			Entry("1", 1, 1),
			Entry("2", 1, 1),
			Entry("3", 3, 2),
			Entry("205", 205, 5),
			Entry("-205", -205, 5),
		)
	})
})

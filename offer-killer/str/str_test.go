package str_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/offer-killer/str"
)

var _ = Describe("Str", func() {
	Context("014-CheckInclusion", func() {
		DescribeTable(
			"CheckInclusion",
			func(s1, s2 string, res bool) {
				Expect(CheckInclusion(s1, s2)).To(Equal(res))
			},
			Entry("normal", "ac", "dgcaf", true),
			Entry("normal", "ab", "dgcaf", false),
		)
	})

	Context("015-FindAnagrams", func() {
		DescribeTable(
			"FindAnagrams",
			func(s1, s2 string, res []int) {
				Expect(FindAnagrams(s1, s2)).To(Equal(res))
			},
			Entry("normal", "ac", "dgcaf", []int{2}),
			Entry("normal", "ab", "dgcaf", []int{}),
		)
	})

	Context("016-LengthOfLongestSubstring", func() {
		DescribeTable(
			"LengthOfLongestSubstring",
			func(s string, res int) {
				Expect(LengthOfLongestSubstring(s)).To(Equal(res))
			},
			Entry("normal", "babcca", 3),
			Entry("normal", "asdfasdcdeaac", 5),
		)
	})
})

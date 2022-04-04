package p100_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/leetcode/common"
	. "github.com/mucean/algorithms-go/leetcode/p100"
)

var _ = Describe("P100", func() {
	Context("LRU", func() {
		var lru LRUCache
		var c int
		JustBeforeEach(func() {
			lru = Constructor(c)
		})

		Context("capacity: 1", func() {
			BeforeEach(func() {
				c = 1
			})

			It("1", func() {
				lru.Put(1, 1)
				Expect(lru.Get(1)).To(Equal(1))
				lru.Put(2, 2)
				Expect(lru.Get(2)).To(Equal(2))
			})
		})
	})

	Context("p103", func() {
		DescribeTable(
			"p103",
			func(nodes []int, zLevelNodes [][]int) {
				Expect(ZigzagLevelOrder(CreateTreeNode(nodes))).To(Equal(zLevelNodes))
			},
			Entry("nil node", []int{}, [][]int(nil)),
			Entry("only root node", []int{1}, [][]int{{1}}),
			Entry("deep nodes: 2", []int{1, 2, 3}, [][]int{{1}, {3, 2}}),
			Entry("deep nodes: 2", []int{3, 9, 20, NilTNode, NilTNode, 15, 7}, [][]int{{3}, {20, 9}, {15, 7}}),
		)
	})
})

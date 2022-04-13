package tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/common"
	. "github.com/mucean/algorithms-go/offer/tree"
)

var _ = Describe("Tree", func() {
	Context("06-ConstructBinaryTreeFromPreMid", func() {
		DescribeTable(
			"ConstructBinaryTreeFromPreMid",
			func(pre, mid, res []int) {
				Expect(TreeLevelScan(ConstructBinaryTreeFromPreMid(pre, mid))).To(Equal(res))
			},
			Entry("normal", []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8}),
			//Entry("normal", CreateTreeNode([]int{1, 2, 3, 4, NilTNode, 5, 6, NilTNode, 7, NilTNode, NilTNode, NilTNode, NilTNode, 8, NilTNode})),
		)
	})

	Context("18-HasSubTree", func() {
		DescribeTable(
			"HasSubTree",
			func(t1, t2 []int, res bool) {
				Expect(HasSubTree(CreateTreeNode(t1), CreateTreeNode(t2))).To(Equal(res))
			},
			Entry("nil, nil", []int{}, []int{}, true),
			Entry("nil, not nil", []int{}, []int{1}, false),
			Entry("not nil, nil", []int{1}, []int{}, true),
			Entry("has", []int{8, 8, 7, 9, 2, NilTNode, NilTNode, NilTNode, NilTNode, 4, 7}, []int{8, 9, 2}, true),
		)
	})

	Context("19-Mirror", func() {
		DescribeTable(
			"Mirror",
			func(nodes, levelScan []int) {
				root := CreateTreeNode(nodes)
				Mirror(root)
				Expect(TreeLevelScan(root)).To(Equal(levelScan))
			},
			Entry("nil", []int{}, []int{}),
			Entry("one node", []int{1}, []int{1}),
			Entry("two node", []int{1, 2}, []int{1, 2}),
			Entry("two node", []int{1, NilTNode, 2}, []int{1, 2}),
			Entry("three node", []int{1, 2, 3}, []int{1, 3, 2}),
			Entry("normal", []int{8, 8, 7, 9, 2, NilTNode, NilTNode, NilTNode, NilTNode, 4, 7}, []int{8, 7, 8, 2, 9, 7, 4}),
		)
	})

	Context("23-LevelOrder", func() {
		DescribeTable(
			"LevelOrder",
			func(nodes []int, res []int) {
				Expect(LevelOrder(CreateTreeNode(nodes))).To(Equal(res))
			},
			Entry("nil", []int(nil), []int{}),
			Entry("one node", []int{1}, []int{1}),
			Entry("normal", []int{1, NilTNode, 2}, []int{1, 2}),
			Entry("normal", []int{1, 2}, []int{1, 2}),
			Entry("normal", []int{1, 2, 3, 4, NilTNode, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}),
		)
	})
})

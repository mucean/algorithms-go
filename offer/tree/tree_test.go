package tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/common"
	. "github.com/mucean/algorithms-go/offer/tree"
)

var _ = Describe("Tree", func() {
	Context("006-ConstructBinaryTreeFromPreMid", func() {
		DescribeTable(
			"ConstructBinaryTreeFromPreMid",
			func(pre, mid, res []int) {
				Expect(TreeLevelScan(ConstructBinaryTreeFromPreMid(pre, mid))).To(Equal(res))
			},
			Entry("normal", []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8}),
			//Entry("normal", CreateTreeNode([]int{1, 2, 3, 4, NilTNode, 5, 6, NilTNode, 7, NilTNode, NilTNode, NilTNode, NilTNode, 8, NilTNode})),
		)
	})
})

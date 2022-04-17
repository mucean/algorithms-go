package common

func CreateListNode(nums []int) *ListNode {
	preNode := &ListNode{}

	end := preNode
	for _, num := range nums {
		end.Next = &ListNode{Val: num}
		end = end.Next
	}

	return preNode.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type ComplexListNode struct {
	Val    int
	Next   *ComplexListNode
	Random *ComplexListNode
}

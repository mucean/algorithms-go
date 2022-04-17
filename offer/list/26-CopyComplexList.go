package list

import (
	. "github.com/mucean/algorithms-go/common"
)

type Node = ComplexListNode

var CopyComplexList = copyComplexList

func copyComplexList(head *Node) *Node {
	if head == nil {
		return nil
	}
	res := &Node{Val: head.Val}
	tmp := res
	headBack := head
	eachMap := map[*Node]*Node{head: tmp}

	for head.Next != nil {
		tmp.Next = &Node{Val: head.Next.Val}
		head = head.Next
		tmp = tmp.Next
		eachMap[head] = tmp
	}

	head = headBack
	for head != nil {
		if head.Random != nil {
			eachMap[head].Random = eachMap[head.Random]
		}
		head = head.Next
	}

	return res
}

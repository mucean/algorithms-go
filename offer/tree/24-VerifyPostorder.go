package tree

import (
	"math"
)

var VerifyPostorder = betterVerifyPostorder

func verifyPostorder(postorder []int) bool {
	l := len(postorder)
	if l == 0 {
		return false
	}
	if l == 1 {
		return true
	}

	root := postorder[l-1]
	pos := l - 2

	for ; pos >= 0 && postorder[pos] > root; pos-- {
	}

	for i := pos; i >= 0; i-- {
		if postorder[i] >= root {
			return false
		}
	}

	if verifyPostorder(postorder[:pos+1]) == false {
		return false
	}

	return verifyPostorder(postorder[pos : l-1])
}

func betterVerifyPostorder(postorder []int) bool {
	var recursive func(postorder []int, min, max int) bool
	recursive = func(postorder []int, min, max int) bool {
		l := len(postorder)
		if l == 0 {
			return true
		}

		root := postorder[l-1]
		if root < min || root > max {
			return false
		}
		pos := l - 2
		if pos < 0 {
			return true
		}
		for ; pos >= 0 && postorder[pos] > root; pos-- {
		}

		if recursive(postorder[:pos+1], min, root) == false {
			return false
		}
		return recursive(postorder[pos+1:l-1], root, max)
	}

	return recursive(postorder, math.MinInt, math.MaxInt)
}

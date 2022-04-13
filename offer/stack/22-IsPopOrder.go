package stack

var IsPopOrder = validateStackSequences

func validateStackSequences(pushed []int, popped []int) bool {
	l := len(pushed)
	lp := len(popped)
	if l != lp {
		return false
	}
	stack := make([]int, 0, len(pushed))

	j := 0
	stackLastPos := -1

	for i := 0; i < l; i++ {
		stack = append(stack, pushed[i])
		stackLastPos++
		if pushed[i] != popped[j] {
			continue
		}

		for ; j < lp && stackLastPos >= 0 && stack[stackLastPos] == popped[j]; j++ {
			stack = stack[:stackLastPos]
			stackLastPos--
		}
	}

	return j == lp
}

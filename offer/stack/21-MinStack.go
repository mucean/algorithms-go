package stack

type minStackCount struct {
	min   int
	count int
}

type MinStack struct {
	stack []int
	min   []minStackCount
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	l := len(s.min)
	if l == 0 || s.min[l-1].min > x {
		s.min = append(s.min, minStackCount{min: x, count: 1})
		return
	}
	if s.min[l-1].min == x {
		s.min[l-1].count++
	}
}

func (s *MinStack) Pop() {
	l := len(s.stack)
	if l == 0 {
		return
	}
	if p := len(s.min) - 1; s.stack[l-1] == s.min[p].min {
		if s.min[p].count == 1 {
			s.min = s.min[:p]
		} else {
			s.min[p].count--
		}
	}
	s.stack = s.stack[:l-1]
}

func (s *MinStack) Top() int {
	l := len(s.stack)
	if l == 0 {
		return 0
	}
	return s.stack[l-1]
}

func (s *MinStack) Min() int {
	l := len(s.min)
	if l == 0 {
		return 0
	}
	return s.min[l-1].min
}

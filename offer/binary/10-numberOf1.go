package binary

var NumberOf1 = numberOf1

func numberOf1(num int) int {
	c := 0
	if num < 0 {
		num = -num
	}
	for num != 0 {
		num &= num - 1
		c++
	}

	return c
}

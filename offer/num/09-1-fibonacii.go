package num

func Fibonacii(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	tmp, pre, ppre := 0, 1, 0
	for i := 2; i <= n; i++ {
		tmp = pre + ppre
		ppre, pre = pre, tmp
	}

	return pre
}

package integer

var CountBits = CountBitsBest

func myCountBits(num int) []int {
	res := make([]int, 0, num+1)
	res = append(res, 0)

	i := 1
	var l int

	for i <= num {
		l = len(res)
		for j := 0; j < l && i <= num; j++ {
			res = append(res, res[j]+1)
			i++
		}
	}

	return res
}

func CountBitsBetter(num int) []int {
	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}

	return res
}

func CountBitsBest(num int) []int {
	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		res[i] = res[i>>1] + i&1
	}

	return res
}

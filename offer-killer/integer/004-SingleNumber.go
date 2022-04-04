package integer

var SingleNumber = mySingleNumber

func mySingleNumber(nums []int) int {
	bitCount := make([]int, 64)

	var i int
	maxLen := 0
	for _, num := range nums {
		i = 0
		if num < 0 {
			num = -num
		}
		for num > 0 {
			if num&1 == 1 {
				bitCount[i]++
			}
			i++
			num = num >> 1
		}
		if maxLen < i {
			maxLen = i
		}
	}
	res := 0
	for i = maxLen - 1; i >= 0; i-- {
		res = (res << 1) + bitCount[i]%3
	}

	return res
}

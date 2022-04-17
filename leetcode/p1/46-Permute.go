package p1

var Permute = permute

func permute(nums []int) [][]int {
	l := len(nums)
	ans := make([][]int, 0)
	var recursive func(p int)
	recursive = func(p int) {
		if p == l {
			return
		}
		if p == l-1 {
			each := make([]int, l)
			copy(each, nums)
			ans = append(ans, each)
			return
		}

		for i := p; i < l; i++ {
			nums[p], nums[i] = nums[i], nums[p]
			recursive(p + 1)
			nums[p], nums[i] = nums[i], nums[p]
		}
	}

	recursive(0)

	return ans
}

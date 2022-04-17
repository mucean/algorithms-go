package p1

import "sort"

var PermuteUnique = permuteUnique

func permuteUnique(nums []int) [][]int {
	l := len(nums)
	ans := make([][]int, 0)
	if l == 0 {
		return ans
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var recursive func(p int)
	recursive = func(p int) {
		if p == l {
			return
		}
		if p == l-1 {
			each := make([]int, l)
			copy(each, nums)
			ans = append(ans, each)
		}

		for i := p; i < l; i++ {
			if i != p && nums[i] == nums[i-1] {
				continue
			}
			// copy
			tmp := nums[i]
			for j := i; j > p; j-- {
				nums[j] = nums[j-1]
			}
			nums[p] = tmp

			recursive(p + 1)
			// back
			tmp = nums[p]
			for j := p; j < i; j++ {
				nums[j] = nums[j+1]
			}
			nums[i] = tmp
		}
	}

	recursive(0)
	return ans
}

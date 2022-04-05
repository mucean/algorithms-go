package array

import "sort"

var ThreeSum = threeSumZero

func threeSumZero(nums []int) [][]int {
	return threeSum(nums, 0)
}

func threeSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	l := len(nums)
	if l < 3 {
		return ans
	}

	sort.Ints(nums)

	for first := 0; first < l; first++ {
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, l-1
		for second < third {
			if second != first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			t := target - nums[first]
			if act := nums[second] + nums[third]; act == t {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
				second++
			} else if act > t {
				third--
			} else {
				second++
			}
		}
	}
	return ans
}

package array

var TwoSum = myTwoSum

func myTwoSum(nums []int, target int) []int {
	l := len(nums)
	if l < 2 {
		return nil
	}

	s, e := 0, l-1
	for s < e {
		if out := nums[s] + nums[e]; out == target {
			return []int{s, e}
		} else if out > target {
			e--
		} else {
			s++
		}
	}
	return nil
}

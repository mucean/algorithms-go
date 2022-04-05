package array

var MinSubArrayLen = minSubArrayLen

func minSubArrayLen(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	p1, p2 := 0, 0

	ans := 0
	sum := nums[0]
	for p2 < l {
		if sum >= target {
			if act := p2 - p1 + 1; ans == 0 || ans > act {
				ans = act
			}
			sum -= nums[p1]
			p1++
		} else {
			sum += nums[p2]
			p2++
		}
	}

	return ans
}

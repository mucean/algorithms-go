package array

var SubarraySumCount = bestSubarraySumCount

func mySubarraySumCount(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	preSum := make([]int, l)
	preSum[0] = nums[0]

	for i := 1; i < l; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	ans := 0
	for j := 0; j < l; j++ {
		for i := 0; i < j; i++ {
			if preSum[j]-preSum[i] == target {
				ans++
			}
		}
		if preSum[j] == target {
			ans++
		}
	}
	return ans
}

func bestSubarraySumCount(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	leftMap := map[int]int{0: 1}
	cnt := 0
	preSum := 0
	for i := 0; i < l; i++ {
		preSum += nums[i]
		cnt += leftMap[preSum-target]
		leftMap[preSum] += 1
	}

	return cnt
}

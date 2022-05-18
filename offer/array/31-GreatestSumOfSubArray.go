package array

var GreatestSumOfSubArray = greatestSumOfSubArray

func greatestSumOfSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	greatestSum := nums[0]
	sum := nums[0]
	for _, num := range nums[1:] {
		if sum <= 0 {
			sum = num
		} else {
			sum += num
		}
		if sum > greatestSum {
			greatestSum = sum
		}
	}
	return greatestSum
}

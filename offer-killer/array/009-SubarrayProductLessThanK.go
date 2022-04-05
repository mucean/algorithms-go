package array

var SubarrayProductLessThanK = subarrayProductLessThanK

func subarrayProductLessThanK(nums []int, k int) [][]int {
	ans := make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ans
	}

	i := 0
	product := 1
	for j := 0; j < l; j++ {
		product *= nums[j]
		for i <= j && product >= k {
			product /= nums[i]
			i++
		}
		if i <= j {
			for b := i; b <= j; b++ {
				ans = append(ans, nums[b:j+1])
			}
		}
	}

	return ans
}

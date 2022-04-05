package array

var PivotIndex = pivotIndex

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	sum := 0
	for i, num := range nums {
		if sum == total-sum-num {
			return i
		}
		sum += num
	}
	return -1
}

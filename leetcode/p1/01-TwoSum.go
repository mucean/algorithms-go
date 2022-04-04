package p1

var TwoSum = twoSum

func twoSum(nums []int, target int) []int {
	l := len(nums)
	if l <= 1 {
		return nil
	}
	numMap := make(map[int]int, l)
	for i, num := range nums {
		rest := target - num
		if p, ok := numMap[rest]; ok {
			return []int{p, i}
		}
		numMap[num] = i
	}
	return nil
}

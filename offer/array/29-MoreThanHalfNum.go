package array

import "math/rand"

var MoreThanHalfNum = moreThanHalfNum

func moreThanHalfNum(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	matchedIndex := l / 2
	start, end := 0, l-1
	var pos int
	for start < end {
		pos = partition(nums, start, end)
		if pos == matchedIndex {
			return nums[pos]
		} else if pos < matchedIndex {
			start = pos
		} else {
			end = pos
		}
	}

	return nums[matchedIndex]
}

func partition(nums []int, start, end int) int {
	if start >= end {
		return start
	}
	pos := rand.Intn(end - start + 1)
	nums[pos], nums[end] = nums[end], nums[pos]
	cmp := nums[end]
	pos = start
	for i := start; i < end; i++ {
		if nums[pos] < cmp {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}

	nums[pos], nums[end] = nums[end], nums[pos]

	return pos
}

// better the number of numbers in the array exceeds half the length of the array.
// so we can increase the count of the same number in the array when we iterate the
// array, when the current number not equal previous counted number, and if the count
// of previous counted number is greater than zero, then decrease it, otherwise make
// the current number as counted number, and the count is one.
// At the end of the iterator, left number is the answer
func better(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	previousNumber := nums[0]
	previousCount := 1

	for _, num := range nums[1:] {
		if num == previousNumber {
			previousCount++
			continue
		}
		if previousCount == 1 {
			previousNumber = num
		} else {
			previousCount--
		}
	}

	return previousNumber
}

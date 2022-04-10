package array

var MinArray = bestMinArray

func myMinArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	} else if l == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	mid := l / 2

	if nums[0] < nums[mid] {
		if nums[mid] < nums[l-1] {
			return nums[0]
		}
		return myMinArray(nums[mid+1:])
	}

	if nums[mid] > nums[l-1] {
		return myMinArray(nums[mid+1:])
	}

	if nums[0] > nums[mid] || nums[mid] < nums[l-1] {
		return myMinArray(nums[0 : mid+1])
	}

	ans := myMinArray(nums[0:mid])
	if n := myMinArray(nums[mid:]); n > ans {
		return ans
	} else {
		return n
	}
}

func bestMinArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}

	return nums[left]
}

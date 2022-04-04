package p200

var FindKthLargest = findKthLargest

func findKthLargest(nums []int, k int) int {
	// 5, 1, 2, 9, 10, 2, 5

	// 1, 5, ....
	// 1, 2, 5, ....
	// 1, 2, 5, 9, 10, ...
	// 1, 2, 2, 5, 10, 9, 5

	var p, pn int
	for {
		l := len(nums)
		if l == 1 {
			return nums[0]
		}

		if k == 1 {
			return findMax(nums)
		}

		if k == l {
			return findMin(nums)
		}

		pn = nums[0]
		for i := 1; i < l; i++ {
			if nums[i] <= pn {
				nums[i], nums[p] = nums[p], nums[i]
				p++
			}
		}

		if s := k - (l - p); s == 0 {
			return pn
		} else if s < 0 {
			nums = nums[p+1:]
		} else {
			nums = nums[:p]
			k = s
		}
		p = 0
	}
}

func findMin(nums []int) int {
	min := nums[0]
	for _, n := range nums[1:] {
		if min > n {
			min = n
		}
	}
	return min
}

func findMax(nums []int) int {
	max := nums[0]
	for _, n := range nums[1:] {
		if max < n {
			max = n
		}
	}
	return max
}

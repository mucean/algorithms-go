package array

var KMinNumbers = kMinNumbers

func kMinNumbers(nums []int, k int) []int {
	l := len(nums)
	if l <= k {
		return nums
	}
	if k <= 0 {
		return nil
	}
	minNumbers := maxHeap(make([]int, k))
	copy(minNumbers, nums)
	minNumbers.init()
	for i := k; i < l; i++ {
		if minNumbers[0] > nums[i] {
			minNumbers[0] = nums[i]
			minNumbers.Down(0)
		}
	}
	return minNumbers
}

type maxHeap []int

func (h maxHeap) init() {
	for i := len(h) / 2; i >= 0; i-- {
		h.Down(i)
	}
}

func (h maxHeap) Down(pos int) {
	l := len(h)
	for pos < l {
		left := pos*2 + 1
		if left >= l {
			break
		}
		if right := left + 1; right < l && h[left] < h[right] {
			left = right
		}
		if h[pos] > h[left] {
			break
		}
		h[pos], h[left] = h[left], h[pos]
		pos = left
	}
}

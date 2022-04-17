package p1

var NextPermutation = nextPermutation

func nextPermutation(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	// find max pos, record change pos = max pos -1
	// find sub max pos: > change pos
	// change
	// s = change pos, e = l -1
	// for s < e => nums[s], nums[e] = nums[e], nums[s]

	maxPos := l - 1
	for ; maxPos > 0; maxPos-- {
		if nums[maxPos-1] < nums[maxPos] {
			break
		}
	}
	var s int
	if maxPos == 0 {
		if nums[0] == nums[1] {
			return
		}
		s = 0
	} else {
		changePos := maxPos - 1
		subMaxPos := maxPos
		for ; subMaxPos < l; subMaxPos++ {
			if nums[subMaxPos] <= nums[changePos] {
				break
			}
		}
		subMaxPos--
		nums[changePos], nums[subMaxPos] = nums[subMaxPos], nums[changePos]
		s = maxPos
	}

	e := l - 1
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
	return
}

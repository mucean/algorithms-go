package array

import (
	"strconv"
)

var FindMaxLength = bestFindMaxLength01

type ZeroOneCount struct {
	Zero int
	One  int
}

func (c ZeroOneCount) String() string {
	return strconv.Itoa(c.Zero) + "-" + strconv.Itoa(c.One)
}

func myFindMaxLength01(nums []int) int {
	preMap := map[string]int{"0-0": -1}
	maxLen := 0
	preSum := ZeroOneCount{}
	l := len(nums)
	findP := 0
	ok := false
	if l == 0 {
		return maxLen
	}

	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			preSum.Zero++
		} else {
			preSum.One++
		}

		if preSum.Zero > preSum.One {
			findP, ok = preMap[ZeroOneCount{Zero: preSum.Zero - preSum.One}.String()]
		} else {
			findP, ok = preMap[ZeroOneCount{Zero: preSum.One - preSum.Zero}.String()]
		}
		if ok && i-findP > maxLen {
			maxLen = i - findP
		}
		preMap[preSum.String()] = i
	}

	return maxLen
}

func bestFindMaxLength01(nums []int) int {
	maxLen := 0
	l := len(nums)
	if l < 2 {
		return maxLen
	}

	preMap := map[int]int{0: -1}
	preSum := 0
	findP := 0
	ok := false
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			preSum += -1
		} else {
			preSum += 1
		}

		findP, ok = preMap[preSum]
		if ok {
			if i-findP > maxLen {
				maxLen = i - findP
			}
		} else {
			preMap[preSum] = i
		}

	}
	return maxLen
}

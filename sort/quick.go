package sort

import (
	"math/rand"
	"time"
)

var randVar = rand.New(rand.NewSource(time.Now().Unix()))

func Partition(nums []int, start, end int) int {
	if start >= end {
		return end
	}

	index := randVar.Intn(end-start+1) + start
	comp := nums[index]
	nums[index], nums[end] = nums[end], nums[index]
	index = start - 1
	for i := start; i < end; i++ {
		if nums[i] <= comp {
			index++
			if i != index {
				nums[i], nums[index] = nums[index], nums[i]
			}
		}
	}
	index++
	nums[index], nums[end] = nums[end], nums[index]

	return index
}

func quick(nums []int, start, end int) {
	i := Partition(nums, start, end)

	if i > start {
		quick(nums, 0, i-1)
	}

	if i < end {
		quick(nums, i+1, end)
	}
}

func Quick(nums []int) {
	quick(nums, 0, len(nums)-1)
}

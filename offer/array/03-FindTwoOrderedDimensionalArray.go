package array

var FindTwoOrderedDimensionalArray = findTwoOrderedDimensionalArray

func findTwoOrderedDimensionalArray(nums [][]int, target int) bool {
	// array nil
	// find: target == min, target == max, target == middle
	// not found: equal find

	row := len(nums)
	if row == 0 {
		return false
	}

	col := len(nums[0])
	if col == 0 {
		return false
	}

	// min, max
	if nums[0][0] == target || nums[row-1][col-1] == target {
		return true
	} else if nums[0][0] > target || nums[row-1][col-1] < target {
		return false
	}

	x, y := row-1, 0

	for x > 0 && y < col {
		if nums[x][y] == target {
			return true
		}

		if nums[x][y] > target {
			x--
		} else {
			y++
		}
	}

	return false
}

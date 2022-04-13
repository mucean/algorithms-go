package array

var SpiralOrder = spiralOrder

func spiralOrder(matrix [][]int) []int {
	rmin, rmax := 0, len(matrix)
	if rmax == 0 {
		return []int{}
	}
	cmin, cmax := 0, len(matrix[0])
	if cmax == 0 {
		return []int{}
	}

	ans := make([]int, 0, rmax*cmax)

	for rmin < rmax && cmin < cmax {
		for i := cmin; i < cmax; i++ {
			ans = append(ans, matrix[rmin][i])
		}
		rmin++
		for i := rmin; i < rmax; i++ {
			ans = append(ans, matrix[i][cmax-1])
		}
		cmax--
		for i := cmax - 1; i >= cmin && rmin < rmax; i-- {
			ans = append(ans, matrix[rmax-1][i])
		}
		rmax--
		for i := rmax - 1; i >= rmin && cmin < cmax; i-- {
			ans = append(ans, matrix[i][cmin])
		}
		cmin++
	}

	return ans
}

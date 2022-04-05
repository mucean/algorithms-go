package array

type NumMatrix [][]int

func (m NumMatrix) SubMatrixSum(row1, col1, row2, col2 int) int {
	return m[row2+1][col2+1] - m[row2+1][col1] - m[row1][col2+1] + m[row1][col1]
}

func NewNumMatrix(nums [][]int) NumMatrix {
	row := len(nums)
	matrix := make(NumMatrix, 0, row+1)
	if row == 0 {
		return matrix
	}
	col := len(nums[0])
	for i := 0; i <= row; i++ {
		matrix = append(matrix, make([]int, col+1))
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1] + nums[i-1][j-1]
		}
	}

	return matrix
}

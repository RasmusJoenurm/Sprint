package sprint

func TransposeMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	numRows := len(matrix)
	numCols := len(matrix[0])
	transposed := make([][]int, numCols)
	for i := 0; i < numCols; i++ {
		transposed[i] = make([]int, numRows)
	}
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			transposed[c][r] = matrix[r][c]
		}
	}
	return transposed
}

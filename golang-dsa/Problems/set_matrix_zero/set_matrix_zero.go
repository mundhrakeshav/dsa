package setmatrixzero

func SetMatrixZero(matrix [][]int) [][]int {
	cols := len(matrix)
	rows := len(matrix[0])
	matrix_temp := make([][]int, rows)
	for i, _ := range matrix {
		matrix_temp[i] = append(matrix_temp[i], matrix[i]...)
	}

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				for a := 0; a < rows; a++ {
					matrix_temp[i][a] = 0
				}
				for a := 0; a < cols; a++ {
					matrix_temp[a][j] = 0
				}
			}
		}
	}
	return matrix_temp
}

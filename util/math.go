package util

func Sum(input []int) int {
	s := 0
	for _, element := range input {
		s = s + element
	}
	return s
}

func Transpose(m [][]rune) [][]rune {
	x := len(m[0])
	y := len(m)

	// Create matrix with new dimensions
	matrix := make([][]rune, x)
	for i := range matrix {
		matrix[i] = make([]rune, y)
	}

	// Transpose matrix
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			matrix[i][j] = m[j][i]
		}
	}

	return matrix
}

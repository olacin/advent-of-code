package util

import (
	"math"
	"sort"
)

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

func Median(s []int) int {
	sort.Ints(s)

	n := len(s)

	if n%2 == 0 {
		return (s[n/2] + s[(n/2)-1]) / 2
	}
	return s[(n-1)/2]
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(numbers ...int) int {
	var max = math.MinInt64
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

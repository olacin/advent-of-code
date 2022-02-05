package day05

import (
	"fmt"
)

type Range struct {
	Start int
	End   int
}

func (r Range) Slice() []int {
	var slice []int

	if r.Start > r.End {
		for i := r.End; i <= r.Start; i++ {
			slice = append(slice, i)
		}
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
	} else {
		for i := r.Start; i <= r.End; i++ {
			slice = append(slice, i)
		}
	}

	return slice
}

type Line struct {
	X Range
	Y Range
}

func (l Line) IsStraight() bool {
	return l.X.Start == l.X.End || l.Y.Start == l.Y.End
}

func (l Line) IsDiagonal() bool {
	xabs := Abs(l.X.End - l.X.Start)
	yabs := Abs(l.Y.End - l.Y.Start)
	return xabs == yabs
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInput(content []string) []Line {
	var lines []Line

	for _, line := range content {
		var x1, x2, y1, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if x2 < x1 {
			x1, x2 = x2, x1
		}
		if y2 < y1 {
			y1, y2 = y2, y1
		}

		r1 := Range{Start: x1, End: x2}
		r2 := Range{Start: y1, End: y2}

		l := Line{r1, r2}

		lines = append(lines, l)
	}

	return lines
}

func NewFloor(lines []Line) [][]int {
	// Find floor size
	var xsize, ysize int
	for _, l := range lines {
		xmax := max(l.X.Start, l.X.End)
		ymax := max(l.Y.Start, l.Y.End)
		if xmax > xsize {
			xsize = xmax
		}
		if ymax > ysize {
			ysize = ymax
		}
	}

	// Init floor
	floor := make([][]int, ysize+1)
	for i := range floor {
		floor[i] = make([]int, xsize+1)
	}

	return floor
}

func PrintFloor(floor [][]int) {
	for i := range floor {
		for j := range floor[i] {
			v := floor[i][j]
			if v == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", v)
			}
		}
		fmt.Println()
	}
}

func CountOverlaps(floor [][]int) int {
	count := 0
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			if floor[i][j] > 1 {
				count += 1
			}
		}
	}
	return count
}

func Filter(lines []Line, diagonals bool) []Line {
	filtered := make([]Line, 0)
	for _, l := range lines {
		if l.IsStraight() || (diagonals && l.IsStraight()) {
			filtered = append(filtered, l)
		}
	}
	return filtered
}

func part1(lines []Line) int {
	floor := NewFloor(lines)

	for _, l := range lines {
		if l.IsStraight() {
			for _, x := range l.X.Slice() {
				for _, y := range l.Y.Slice() {
					floor[y][x] += 1
				}
			}
		}
	}

	return CountOverlaps(floor)
}

func part2(lines []Line) [][]int {
	floor := NewFloor(lines)

	for _, l := range lines[:2] {
		if l.IsDiagonal() {
			fmt.Println(l, "is a diagonal")
			for _, x := range l.X.Slice() {
				for _, y := range l.Y.Slice() {
					fmt.Printf("diag (%d, %d)\n", x, y)
					floor[y][x] += 1
				}
			}
		}
	}

	return floor
}

// Part 1: 4993
// floor := part1(lines)
// PrintFloor(floor)
// fmt.Printf("Part 1: Overlapping points: %d\n", CountOverlaps(floor))

// Part 2:
// floor := part2(lines)
// PrintFloor(floor)
// fmt.Printf("Part 2: Overlapping points: %d\n", CountOverlaps(floor))

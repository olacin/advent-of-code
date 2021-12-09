package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(path string) [][]int {
	content, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(content), "\n")

	floor := make([][]int, 0)

	for _, l := range lines {
		if l == "" {
			continue
		}
		row := make([]int, 0)
		for _, r := range l {
			n, _ := strconv.Atoi(string(r))
			row = append(row, n)
		}
		floor = append(floor, row)
	}

	return floor
}

func IsLow(value int, neighbors []int) bool {
	for _, n := range neighbors {
		if value < n {
			continue
		}
		return false
	}
	return true
}

func Neighbors(floor [][]int, i int, j int) [][]int {
	height := len(floor)
	width := len(floor[0])

	neighbors := make([][]int, 0)
	// Check top neighbor
	if i > 0 {
		// neighbors = append(neighbors, floor[i-1][j])
		neighbors = append(neighbors, []int{i - 1, j})
	}
	// Check bottom neighbor
	if i < height-1 {
		// neighbors = append(neighbors, floor[i+1][j])
		neighbors = append(neighbors, []int{i + 1, j})
	}
	// Check left neighbor
	if j > 0 {
		// neighbors = append(neighbors, floor[i][j-1])
		neighbors = append(neighbors, []int{i, j - 1})
	}
	// Check right neighbor
	if j < width-1 {
		// neighbors = append(neighbors, floor[i][j+1])
		neighbors = append(neighbors, []int{i, j + 1})
	}

	return neighbors
}

func LowPoints(floor [][]int) [][]int {
	height := len(floor)
	width := len(floor[0])

	points := make([][]int, 0)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// Get neighbors values
			neighbors := make([]int, 0)
			for _, pos := range Neighbors(floor, i, j) {
				x, y := pos[0], pos[1]
				neighbors = append(neighbors, floor[x][y])
			}

			// Check if it is a low point
			if IsLow(floor[i][j], neighbors) {
				points = append(points, []int{i, j})
			}
		}
	}

	return points
}

func DFS(floor [][]int, i int, j int) int {
	size := 0
	value := floor[i][j]
	// Mark as visited
	floor[i][j] = -1
	for _, n := range Neighbors(floor, i, j) {
		x, y := n[0], n[1]
		if floor[x][y] == -1 {
			continue
		}
		if floor[x][y] > value && floor[x][y] != 9 {
			size += DFS(floor, x, y) + 1
		}
	}

	return size
}

func Part1(floor [][]int) int {
	points := LowPoints(floor)
	risk := 0

	for _, point := range points {
		i, j := point[0], point[1]
		risk += floor[i][j] + 1
	}

	return risk
}

func Part2(floor [][]int) int {
	points := LowPoints(floor)
	basins := make([]int, 0)

	// Compute basins
	for _, point := range points {
		i, j := point[0], point[1]
		basins = append(basins, DFS(floor, i, j)+1)
	}

	// Reverse sort basins sizes
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return basins[0] * basins[1] * basins[2]
}

func main() {
	path := os.Args[1]
	floor := parseInput(path)
	fmt.Println("Part 1: Risk Level: ", Part1(floor))
	fmt.Println("Part 2: ", Part2(floor))
}

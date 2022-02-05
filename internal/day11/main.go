package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Octopus struct {
	Energy int
	Flashed bool
}

func flash(octopuses [][]Octopus, i, j int) int {
	count := 1
	positions := []int{-1, 0, 1}

	for _, x := range positions {
		for _, y := range positions {
			i, j := i+x, j+y
			if i >= 0 && i < len(octopuses) && j >= 0 && j < len(octopuses[0]) && !octopuses[i][j].Flashed {
				octopuses[i][j].Energy += 1
				if octopuses[i][j].Energy > 9 {
					count += flash(octopuses, i, j)
				}
			}
		}
	}

	return count
}

func parseInput(path string) [][]Octopus {
	_content, _ := ioutil.ReadFile(path)
	content := strings.Split(string(_content), "\n")

	octopuses := make([][]Octopus, 0)

	for _, line := range content {
		if line == "" {
			continue
		}
		row := make([]Octopus, 0)
		for _, level := range line {
			n, _ := strconv.Atoi(string(level))
			row = append(row, Octopus{Energy: n, Flashed: false})
		}
		octopuses = append(octopuses, row)
	}

	return octopuses
}


func PrintOctopuses(octopuses [][]Octopus) {
	for i := range octopuses {
		for j := range octopuses[i] {
			fmt.Printf("%d ", octopuses[i][j].Energy)
		}
		fmt.Println()
	}
}


func increment(octopuses [][]Octopus) [][]Octopus {
	for i := range octopuses {
		for j := range octopuses[i] {
			octopuses[i][j].Energy += 1
		}
	}
	return octopuses
}

func main() {
	path := os.Args[1]
	octopuses := parseInput(path)
	PrintOctopuses(octopuses)
	octopuses = increment(octopuses)
	// count := 0
	// for i := 0; i < len(octopuses); i++ {
	// 	for j := 0; j < len(octopuses[i]); j++ {
	// 		count += flash(octopuses, i, j)
	// 	}
	// }
	fmt.Println()
	flash(octopuses, 0, 0)
	PrintOctopuses(octopuses)
	// fmt.Println("Flash count: ", count)
}

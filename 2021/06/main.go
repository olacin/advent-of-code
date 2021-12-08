package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(path string) []int {
	content, _ := ioutil.ReadFile(path)
	line := strings.Split(string(content), "\n")[0]

	fishes := make([]int, 0)

	for _, s := range strings.Split(line, ",") {
		n, _ := strconv.Atoi(s)
		fishes = append(fishes, n)
	}

	return fishes
}

func memoize() func(int, int) int {
	type args struct {
		days int
		fish int
	}

	memo := make(map[args]int)

	var function func(int, int) int
	function = func(days int, fish int) int {
		days = days - (fish + 1)
		if days < 0 {
			return 1
		}
		if _, ok := memo[args{days, fish}]; !ok {
			memo[args{days, fish}] = function(days, 6) + function(days, 8)
		}
		return memo[args{days, fish}]
	}
	return function
}

func compute(days int, fishes []int) int {
	total := 0
	fn := memoize()
	for _, f := range fishes {
		total += fn(days, f)
	}
	return total
}

func main() {
	fishes := parseInput("input.txt")

	// Part 1: 351188
	fmt.Printf("Number of fishes after 80 days: %d\n", compute(80, fishes))
	// Part 2:
	fmt.Printf("Number of fishes after 256 days: %d\n", compute(256, fishes))
}

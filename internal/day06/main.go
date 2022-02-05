package day06

import (
	"strings"

	"github.com/olacin/advent-of-code/util"
)

func parseInput(lines []string) []int {
	fishes := make([]int, 0)

	for _, s := range strings.Split(lines[0], ",") {
		fishes = append(fishes, util.Atoi(s))
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

// Part 1: 351188
// Part 2:

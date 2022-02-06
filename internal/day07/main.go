package day07

import (
	"math"
	"strings"

	"github.com/olacin/advent-of-code/util"
)

func parseInput(lines []string) []int {
	crabs := make([]int, 0)

	for _, s := range strings.Split(lines[0], ",") {
		crabs = append(crabs, util.Atoi(s))
	}

	return crabs
}

func Series(n int) int {
	return (n * (n + 1)) / 2
}

func part1(crabs []int) int {
	med := util.Median(crabs)
	fuel := 0
	for _, c := range crabs {
		fuel += util.Abs(c - med)
	}
	return fuel
}

func part2(crabs []int) int {
	least := math.MaxInt64
	for med := 0; med < util.Max(crabs...); med++ {
		fuel := 0
		for _, c := range crabs {
			d := util.Abs(c - med)
			fuel += Series(d)
		}
		if fuel < least {
			least = fuel
		}
	}
	return least
}

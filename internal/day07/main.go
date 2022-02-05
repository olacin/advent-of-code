package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseInput(path string) []int {
	content, _ := ioutil.ReadFile(path)
	line := strings.Split(string(content), "\n")[0]

	crabs := make([]int, 0)

	for _, s := range strings.Split(line, ",") {
		n, _ := strconv.Atoi(s)
		crabs = append(crabs, n)
	}

	return crabs
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

func Series(n int) int {
	return (n * (n + 1)) / 2
}

func part1(crabs []int) int {
	med := Median(crabs)
	fuel := 0
	for _, c := range crabs {
		fuel += Abs(c - med)
	}
	return fuel
}

func part2(crabs []int) int {
	least := math.MaxInt64
	for med := 0; med < Max(crabs...); med++ {
		fuel := 0
		for _, c := range crabs {
			d := Abs(c - med)
			fuel += Series(d)
		}
		if fuel < least {
			least = fuel
		}
	}
	return least
}

func main() {
	crabs := parseInput("input.txt")
	fmt.Printf("Part 1: Fuel consumption: %d\n", part1(crabs))
	fmt.Printf("Part 2: Fuel consumption: %d\n", part2(crabs))
}

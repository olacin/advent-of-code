package day01

import "github.com/olacin/advent-of-code/util"

func part1(measurements []int) int {
	count := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i-1] < measurements[i] {
			count = count + 1
		}
	}
	return count
}

func part2(measurements []int) int {
	count := 0
	var _measurements []int

	// Group measurements in tuples
	for i := 0; i < len(measurements)-2; i++ {
		_measurements = append(_measurements, util.Sum(measurements[i:i+3]))
	}

	// Count increased measurements
	for i := 1; i < len(_measurements); i++ {
		if _measurements[i-1] < _measurements[i] {
			count = count + 1
		}
	}

	return count
}

package day03

import (
	"github.com/olacin/advent-of-code/util"
	"strconv"
)

func part1(numbers [][]rune) int {
	var g, e string
	numbers = util.Transpose(numbers)
	for _, v := range numbers {
		if util.RuneOccurrences(v, '1') > util.RuneOccurrences(v, '0') {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}

	gamma, _ := strconv.ParseInt(g, 2, 64)
	epsilon, _ := strconv.ParseInt(e, 2, 64)

	return int(gamma * epsilon)
}

func filter(numbers [][]rune, r rune, index int) [][]rune {
	_numbers := make([][]rune, 0)
	for _, n := range numbers {
		if n[index] == r {
			_numbers = append(_numbers, n)
		}
	}
	return _numbers
}

func rating(numbers [][]rune, index int, reversed bool) string {
	if len(numbers) == 1 {
		return string(numbers[0])
	} else {
		column := util.Transpose(numbers)[index]

		nbOnes := util.RuneOccurrences(column, '1')
		nbZeroes := util.RuneOccurrences(column, '0')

		if reversed {
			if nbZeroes <= nbOnes {
				return rating(filter(numbers, '0', index), index+1, reversed)
			} else {
				return rating(filter(numbers, '1', index), index+1, reversed)
			}
		} else {
			if nbOnes >= nbZeroes {
				return rating(filter(numbers, '1', index), index+1, reversed)
			} else {
				return rating(filter(numbers, '0', index), index+1, reversed)
			}
		}
	}
}

func part2(numbers [][]rune) int {
	o := rating(numbers, 0, false)
	c := rating(numbers, 0, true)

	oxygen, _ := strconv.ParseInt(o, 2, 64)
	co2, _ := strconv.ParseInt(c, 2, 64)

	return int(oxygen * co2)
}

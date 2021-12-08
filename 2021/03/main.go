package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput(path string) ([][]rune, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		numbers = append(numbers, runes)
	}

	return numbers, scanner.Err()
}

func transpose(m [][]rune) [][]rune {
	x := len(m[0])
	y := len(m)

	// Create matrix with new dimensions
	_m := make([][]rune, x)
	for i := range _m {
		_m[i] = make([]rune, y)
	}

	// Transpose matrix
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			_m[i][j] = m[j][i]
		}
	}

	return _m
}

func count(s []rune, r rune) int {
	c := 0
	for _, _r := range s {
		if _r == r {
			c += 1
		}
	}
	return c
}

func part1(numbers [][]rune) int64 {
	var g, e string
	for _, v := range numbers {
		if count(v, '1') > count(v, '0') {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}

	gamma, _ := strconv.ParseInt(g, 2, 64)
	epsilon, _ := strconv.ParseInt(e, 2, 64)

	return gamma * epsilon
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
		column := transpose(numbers)[index]

		nbOnes := count(column, '1')
		nbZeroes := count(column, '0')

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

func part2(numbers [][]rune) int64 {
	o := rating(numbers, 0, false)
	c := rating(numbers, 0, true)

	oxygen, _ := strconv.ParseInt(o, 2, 64)
	co2, _ := strconv.ParseInt(c, 2, 64)

	return oxygen * co2
}

func main() {
	numbers, err := parseInput("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Part 1: 4147524
	fmt.Printf("Part 1: Power consumption: %d\n", part1(transpose(numbers)))

	// Part 2: 3570354
	fmt.Printf("Part 2: Life support rating: %d\n", part2(numbers))
}

package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func readlines(contents string) []string {
	lines := make([]string, 0)
	for _, line := range strings.Split(contents, "\n") {
		// strings.Split can return empty lines
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

func processLine(line string) (int, error) {
	numbers := make([]rune, 0)
	for _, chr := range line {
		if unicode.IsDigit(chr) {
			numbers = append(numbers, chr)
		}
	}

	first := numbers[0]
	last := numbers[len(numbers)-1]

	return strconv.Atoi(fmt.Sprintf("%c%c", first, last))
}

func sum(numbers []int) int {
	n := 0
	for _, nb := range numbers {
		n += nb
	}
	return n
}

func part1(lines []string) int {
	numbers := make([]int, len(lines))

	for _, line := range lines {
		nb, _ := processLine(line)
		numbers = append(numbers, nb)
	}

	return sum(numbers)
}

func processLine2(line string) (int, error) {
	hm := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	numbers := make([]rune, 0)
	for i, c := range line {
		if unicode.IsDigit(c) {
			numbers = append(numbers, c)
		}
		for str, digit := range hm {
			if strings.HasPrefix(line[i:], str) {
				numbers = append(numbers, digit)
			}
		}
	}

	first := numbers[0]
	last := numbers[len(numbers)-1]

	return strconv.Atoi(fmt.Sprintf("%c%c", first, last))
}

func part2(lines []string) int {
	numbers := make([]int, len(lines))

	for _, line := range lines {
		nb, err := processLine2(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, nb)
	}

	return sum(numbers)
}

func main() {
	lines := readlines(input)
	s := part1(lines)
	fmt.Println("part1: sum is", s)
	s2 := part2(lines)
	fmt.Println("part2: sum is", s2)
}

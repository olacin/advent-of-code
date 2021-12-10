package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var mapping = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var scoresCompletion = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func parseInput(path string) []string {
	content, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(content), "\n")
	return lines[:len(lines)-1]
}

func IsIn(r rune, s string) bool {
	for _, _rune := range s {
		if _rune == r {
			return true
		}
	}
	return false
}

func median(s []int) int {
	sort.Ints(s)

	n := len(s)

	if n%2 == 0 {
		return (s[n/2] + s[(n/2)-1]) / 2
	}
	return s[(n-1)/2]
}

func compute(lines []string) (int, int) {
	score := 0
	scoreCompletion := make([]int, 0)

	for _, line := range lines {
		s := make([]rune, 0)
		var corrupted bool
		var last rune

		for _, r := range line {
			// Opening rune
			if IsIn(r, "([{<") {
				s = append(s, r)
			} else {
				last = s[len(s)-1]
				s = s[:len(s)-1]
				if last != mapping[r] {
					score += scores[r]
					corrupted = true
					break
				}
			}
		}

		if !corrupted {
			count := 0
			for len(s) > 0 {
				r := s[len(s)-1]
				s = s[:len(s)-1]
				for k, v := range mapping {
					if r == v {
						count = count*5 + scoresCompletion[k]
						break
					}
				}
			}
			scoreCompletion = append(scoreCompletion, count)
		}
	}

	return score, median(scoreCompletion)
}

func main() {
	path := os.Args[1]
	lines := parseInput(path)
	part1, part2 := compute(lines)
	fmt.Printf("part 1 - score=%d\n", part1)
	fmt.Printf("part 2 - score=%d\n", part2)
}

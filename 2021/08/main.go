package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var numbers = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func parseInput(path string) [][]string {
	content, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(content), "\n")

	entries := make([][]string, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		entry := make([]string, 0)
		for _, field := range strings.Fields(line) {
			if field == "|" {
				continue
			}
			entry = append(entry, field)
		}
		entries = append(entries, entry)
	}

	return entries
}

func segments(digit string) int {
	counter := make(map[rune]int)
	for _, r := range digit {
		counter[r] += 1
	}
	return len(counter)
}

func Part1(entries [][]string) int {
	total := 0
	counter := make(map[int]int)

	for _, entry := range entries {
		for _, digit := range entry[10:] {
			if number, ok := numbers[segments(digit)]; ok {
				counter[number] += 1
			}
		}
	}

	for _, v := range counter {
		total += v
	}

	return total
}

func main() {
	path := os.Args[1]
	entries := parseInput(path)
	fmt.Printf("part 1 - total=%d\n", Part1(entries))
}

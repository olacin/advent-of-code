package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Part1(entries [][]string) int {
	total := 0
	counter := make(map[int]int)

	for _, entry := range entries {
		for _, digit := range entry[10:] {
			switch segments(digit) {
			case 2:
				counter[1] += 1
			case 4:
				counter[4] += 1
			case 3:
				counter[7] += 1
			case 7:
				counter[8] += 1
			}
		}
	}

	for _, v := range counter {
		total += v
	}

	return total
}

func Part2(entries [][]string) int {
	/* Make a mapper with entries from ten first values */
	/* Once done, compare each output value with mapper values */
	total := 0
	// mapper := make(map[int]string)

	// Compare values by sorting Strings and using strings.Contains()

	for _, entry := range entries {
		// var output string
		for _, digit := range entry[:10] {
			nbSegments := segments(digit)
			fmt.Println(nbSegments)

			// if number, ok := numbers[nbSegments]; ok {
			// 	// Unique nb of segments: 1,4,7,8
			// 	n := strconv.Itoa(number)
			// 	output += string(n)
			// } else if nbSegments == 5 {
			// 	// 5 segments: 2,3,5
			// } else {
			// 	// 6 segments: 0,6,9
			// }
		}
	}

	return total
}

func main() {
	path := os.Args[1]
	entries := parseInput(path)
	fmt.Printf("part 1 - total=%d\n", Part1(entries))
	fmt.Printf("part 2 - total=%d\n", Part2(entries))
}

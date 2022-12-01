package day08

import (
	"sort"
	"strconv"
	"strings"
)

func parseInput(lines []string) [][]string {
	entries := make([][]string, 0)

	for _, line := range lines {
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

func sortEntry(entry []string) []string {
	sort.Slice(entry, func(i, j int) bool {
		return len(entry[i]) < len(entry[j])
	})
	return entry
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func contains(s string, s2 string) string {
	set := make([]rune, 0)
	for _, el := range s {
		if strings.Contains(s2, string(el)) {
			set = append(set, el)
		}
	}
	return string(set)
}

func translate(entry []string, mapper map[int]string) string {
	var decoded string
	for _, field := range entry {
		for k, v := range mapper {
			if field == v {
				decoded += strconv.Itoa(k)
			}
		}
	}
	return decoded
}

func part1(entries [][]string) int {
	total := 0
	counter := make(map[int]int)

	for _, entry := range entries {
		for _, digit := range entry[10:] {
			switch len(digit) {
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

func SortEntries(entries [][]string) [][]string {
	_entries := make([][]string, 0)

	for _, entry := range entries {
		e := make([]string, 0)
		for _, field := range entry {
			e = append(e, sortString(field))
		}
		_entries = append(_entries, e)
	}

	return _entries
}

func part2(entries [][]string) int {
	total := 0

	entries = SortEntries(entries)

	for _, entry := range entries {
		mapper := make(map[int]string)
		for _, digit := range sortEntry(entry[:10]) {
			digit = sortString(digit)
			switch len(digit) {
			case 2:
				mapper[1] = digit
			case 3:
				mapper[7] = digit
			case 4:
				mapper[4] = digit
			case 5:
				if contains(digit, mapper[7]) == mapper[7] {
					mapper[3] = digit
				} else if len(contains(digit, mapper[4])) == 3 {
					mapper[5] = digit
				} else {
					mapper[2] = digit
				}
			case 6:
				if contains(digit, mapper[4]) == mapper[4] {
					mapper[9] = digit
				} else if contains(digit, mapper[7]) == mapper[7] {
					mapper[0] = digit
				} else {
					mapper[6] = digit
				}
			case 7:
				mapper[8] = digit
			}
		}
		n, _ := strconv.Atoi(translate(entry[10:], mapper))
		total += n
	}

	return total
}

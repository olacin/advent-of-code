package util

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadStrings(path string, ignoreEmpty bool) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	raw := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, line := range raw {
		if ignoreEmpty && line == "" {
			continue
		}
		lines = append(lines, line)
	}

	return lines
}

func ReadInts(path string) []int {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	raw := strings.Split(string(content), "\n")

	lines := make([]int, 0)
	for _, line := range raw {
		if line == "" {
			continue
		}
		lines = append(lines, Atoi(line))
	}

	return lines
}

func ReadRunes(path string, ignoreEmpty bool) [][]rune {
	lines := ReadStrings(path, ignoreEmpty)

	runes := make([][]rune, 0)
	for _, line := range lines {
		runes = append(runes, []rune(line))
	}

	return runes
}

func RuneOccurrences(s []rune, r rune) int {
	c := 0
	for _, _r := range s {
		if _r == r {
			c += 1
		}
	}
	return c
}

func Atoi(s string) int {
	integer, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return integer
}

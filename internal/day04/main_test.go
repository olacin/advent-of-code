package day04

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", false)
	numbers, boards := parseInput(input)
	assert.Equal(t, part1(numbers, boards), 16674)
}

func TestPart2(t *testing.T) {
	input := util.ReadStrings("input.txt", false)
	numbers, boards := parseInput(input)
	assert.Equal(t, part2(numbers, boards), 7075)
}
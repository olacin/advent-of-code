package day09

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	floor := parseInput(input)
	assert.Equal(t, part1(floor), 504)
}

func TestPart2(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	floor := parseInput(input)
	assert.Equal(t, part2(floor), 1558722)
}

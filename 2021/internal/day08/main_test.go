package day08

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	lines := parseInput(input)
	assert.Equal(t, part1(lines), 412)
}

func TestPart2(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	lines := parseInput(input)
	assert.Equal(t, part2(lines), 978171)
}

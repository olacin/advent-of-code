package day07

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	crabs := parseInput(input)
	assert.Equal(t, part1(crabs), 355592)
}

func TestPart2(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	crabs := parseInput(input)
	assert.Equal(t, part2(crabs), 101618069)
}

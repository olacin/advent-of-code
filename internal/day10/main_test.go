package day10

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	part1, _ := compute(input)
	assert.Equal(t, part1, 243939)
}

func TestPart2(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	_, part2 := compute(input)
	assert.Equal(t, part2, 2421222841)
}

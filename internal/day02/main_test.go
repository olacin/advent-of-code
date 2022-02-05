package day02

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	moves := ParseMoves(input)
	assert.Equal(t, part1(moves), 1762050)
}

func TestPart2(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	moves := ParseMoves(input)
	assert.Equal(t, part2(moves), 1855892637)
}

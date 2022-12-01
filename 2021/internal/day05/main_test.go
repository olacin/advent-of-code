package day05

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	lines := parseInput(input)
	assert.Equal(t, part1(lines), 4993)
}

func TestPart2(t *testing.T) {
	t.Skip("Not implemented yet")
}

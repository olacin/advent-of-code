package day03

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadRunes("input.txt", true)
	assert.Equal(t, part1(input), 4147524)
}

func TestPart2(t *testing.T) {
	input := util.ReadRunes("input.txt", true)
	assert.Equal(t, part2(input), 3570354)
}

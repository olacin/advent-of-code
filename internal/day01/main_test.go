package day01

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	measurements := util.ReadInts("input.txt")
	assert.Equal(t, part1(measurements), 1393)
}

func TestPart2(t *testing.T) {
	measurements := util.ReadInts("input.txt")
	assert.Equal(t, part2(measurements), 1359)
}

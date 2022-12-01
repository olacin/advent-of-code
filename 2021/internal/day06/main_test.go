package day06

import (
	"testing"

	"github.com/olacin/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	lines := parseInput(input)
	assert.Equal(t, compute(80, lines), 351188)
}

func TestPart2(t *testing.T) {
	input := util.ReadStrings("input.txt", true)
	lines := parseInput(input)
	assert.Equal(t, compute(256, lines), 1595779846729)
}

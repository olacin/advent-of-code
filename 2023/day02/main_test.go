package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var test_input string

func TestPart1(t *testing.T) {
	expected := 8
	games := readlines(test_input)
	s := part1(games)
	if s != expected {
		t.Logf("expected to get %d, got %d", expected, s)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	expected := 2286
	games := readlines(test_input)
	s := part2(games)
	if s != expected {
		t.Logf("expected to get %d, got %d", expected, s)
		t.Fail()
	}
}

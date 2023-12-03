package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var test_input_pt1 string

//go:embed input_test_2.txt
var test_input_pt2 string

func TestPart1(t *testing.T) {
	expected := 142
	lines := readlines(test_input_pt1)
	s := part1(lines)
	if s != expected {
		t.Logf("expected to get %d, got %d", expected, s)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	expected := 281
	lines := readlines(test_input_pt2)
	s := part2(lines)
	if s != expected {
		t.Logf("expected to get %d, got %d", expected, s)
		t.Fail()
	}
}

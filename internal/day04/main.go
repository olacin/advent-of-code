package day04

import (
	"strings"

	"github.com/olacin/advent-of-code/util"
)

type Number struct {
	Value  int
	Marked bool
}

type Board [][]Number

func (b Board) Mark(number int) {
	for i, row := range b {
		for j := range row {
			if b[i][j].Value == number {
				b[i][j].Marked = true
			}
		}
	}
}

func (b Board) Score(number int) int {
	s := 0
	for i, row := range b {
		for j := range row {
			if !b[i][j].Marked {
				s += b[i][j].Value
			}
		}
	}
	return s * number
}

func (b Board) HasWon() bool {
	// Check rows
	for i := 0; i < len(b); i++ {
		won := true
		for j := 0; j < len(b[i]); j++ {
			if !(b[i][j].Marked) {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}

	// Check columns
	for i := 0; i < len(b); i++ {
		won := true
		for j := 0; j < len(b[i]); j++ {
			if !(b[j][i].Marked) {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}

	return false
}

func parseInput(lines []string) ([]int, []Board) {
	var draw []int
	var boards []Board
	var board Board

	// Parse draw
	for _, s := range strings.Split(lines[0], ",") {
		draw = append(draw, util.Atoi(s))
	}

	// Parse boards
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			boards = append(boards, board)
			board = make(Board, 0)
			continue
		}

		row := make([]Number, 0)
		for _, s := range strings.Fields(line) {
			row = append(row, Number{util.Atoi(s), false})
		}
		board = append(board, row)
	}
	boards = append(boards, board)

	return draw, boards
}

func markBoards(number int, boards []Board) {
	for _, b := range boards {
		b.Mark(number)
	}
}

func check(boards []Board) (int, Board) {
	for i, b := range boards {
		if b.HasWon() {
			return i, b
		}
	}
	return -1, nil
}

func part1(numbers []int, boards []Board) int {
	var score int

	for _, n := range numbers {
		markBoards(n, boards)
		_, b := check(boards)
		if b != nil {
			score = b.Score(n)
			break
		}
	}
	return score
}

func part2(numbers []int, boards []Board) int {
	var score int

	for len(boards) > 1 {
		for _, n := range numbers {
			markBoards(n, boards)
			i, b := check(boards)
			if b != nil {
				score = b.Score(n)
				boards = append(boards[:i], boards[i+1:]...)
				break
			}
		}
	}
	return score
}

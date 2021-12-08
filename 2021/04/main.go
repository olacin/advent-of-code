package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	Value  int
	Marked bool
}

type Board [][]Number

func (b Board) Print() {
	for i, row := range b {
		for j := range row {
			if b[i][j].Marked {
				color.Set(color.FgYellow)
				fmt.Printf("%-2d ", b[i][j].Value)
				color.Unset()
			} else {
				fmt.Printf("%-2d ", b[i][j].Value)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

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

func parseInput(path string) ([]int, []Board, error) {
	content, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(content), "\n")

	var draw []int
	var boards []Board
	var board Board

	// Parse draw
	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		draw = append(draw, n)
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
			n, _ := strconv.Atoi(s)
			row = append(row, Number{n, false})
		}
		board = append(board, row)
	}
	boards = append(boards, board)

	return draw, boards, nil
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

func main() {
	numbers, boards, err := parseInput("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Part 1: 16674
	for _, n := range numbers {
		markBoards(n, boards)
		_, b := check(boards)
		if b != nil {
			b.Print()
			fmt.Printf("Part 1: Winning board score: %d\n", b.Score(n))
			break
		}
	}

	fmt.Println()

	// Part 2: 7075
	var winner Board
	var score int

	for len(boards) > 1 {
		for _, n := range numbers {
			markBoards(n, boards)
			i, b := check(boards)
			if b != nil {
				winner = b
				score = b.Score(n)
				boards = append(boards[:i], boards[i+1:]...)
				break
			}
		}
	}
	winner.Print()
	fmt.Printf("Part 2: Winning board score: %d\n", score)
}

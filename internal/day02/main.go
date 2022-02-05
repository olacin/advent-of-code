package day02

import (
	"github.com/olacin/advent-of-code/util"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

type Move struct {
	Direction string
	Amount    int
}

// Used for part 1
func (p *Position) Update(m Move) {
	switch m.Direction {
	case "forward":
		p.Horizontal += m.Amount
	case "down":
		p.Depth += m.Amount
	case "up":
		p.Depth -= m.Amount
	}
}

// Used for part 2
func (p *Position) UpdateB(m Move) {
	switch m.Direction {
	case "forward":
		p.Horizontal += m.Amount
		p.Depth += p.Aim * m.Amount
	case "down":
		p.Aim += m.Amount
	case "up":
		p.Aim -= m.Amount
	}
}

func ParseMoves(lines []string) []Move {
	var moves []Move

	for _, line := range lines {
		s := strings.Split(line, " ")

		m := Move{Direction: s[0], Amount: util.Atoi(s[1])}
		moves = append(moves, m)
	}

	return moves
}

func part1(moves []Move) int {
	pos := Position{0, 0, 0}
	for _, m := range moves {
		pos.Update(m)
	}
	return pos.Horizontal * pos.Depth
}

func part2(moves []Move) int {
	pos := Position{0, 0, 0}
	for _, m := range moves {
		pos.UpdateB(m)
	}
	return pos.Horizontal * pos.Depth
}

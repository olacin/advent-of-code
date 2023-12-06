package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Game struct {
	ID   int
	Sets []map[string]int
}

func readlines(data string) []Game {
	games := make([]Game, 0)
	for ln, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		games = append(games, parseGame(ln+1, line))
	}
	return games
}

func parseGame(id int, line string) Game {
	_, s, _ := strings.Cut(line, ":")
	gameSets := strings.Split(s, ";")
	sets := make([]map[string]int, len(gameSets))
	for i, pt := range gameSets {
		sets[i] = parseSet(pt)
	}
	return Game{
		ID:   id,
		Sets: sets,
	}
}

func parseSet(set string) map[string]int {
	fields := strings.Fields(set)
	s := make(map[string]int)
	for i := 0; i < len(fields)-1; i = i + 2 {
		nb, _ := strconv.Atoi(fields[i])
		color := strings.TrimSuffix(fields[i+1], ",")
		s[color] = nb
	}
	return s
}

func isPossible(game Game, bag map[string]int) bool {
	for _, set := range game.Sets {
		for color, nb := range set {
			if bag[color] < nb {
				return false
			}
		}
	}
	return true
}

func part1(games []Game) int {
	total := 0
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, game := range games {
		if isPossible(game, bag) {
			total += game.ID
		}
	}
	return total
}

func minRequired(game Game) map[string]int {
	minimums := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range game.Sets {
		for color, nb := range set {
			minimums[color] = max(minimums[color], nb)
		}
	}
	return minimums
}

func part2(games []Game) int {
	total := 0
	for _, game := range games {
		mins := minRequired(game)
		power := mins["red"] * mins["green"] * mins["blue"]
		total += power
	}
	return total
}

func main() {
	games := readlines(input)
	s1 := part1(games)
	fmt.Println("Part 1: ", s1)
	s2 := part2(games)
	fmt.Println("Part 2: ", s2)
}

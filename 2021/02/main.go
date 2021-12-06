package main


import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Position struct {
  Horizontal int
  Depth int
  Aim int
}

type Move struct {
  Direction string
  Amount int
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
func (p *Position) Update2(m Move) {
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

func parseInput(path string) ([]Move, error) {
  // Open file
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var moves []Move

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    s := strings.Split(scanner.Text(), " ")

    amount, err := strconv.Atoi(s[1])
    if err != nil {
      return nil, err
    }

    m := Move{Direction: s[0], Amount: amount}
    moves = append(moves, m)
  }

  return moves, scanner.Err()
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
    pos.Update2(m)
  }
  return pos.Horizontal * pos.Depth
}

func main() {
  moves, err := parseInput("input.txt")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // Part 1: 1762050
  fmt.Printf("Part 1: Position X * Y: %d\n", part1(moves))

  // Part 2: 1855892637
  fmt.Printf("Part 2: Position X * Y: %d\n", part2(moves))
}

package main


import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func readMeasurements(path string) ([]int, error) {
  // Open file
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  // Scan measurements
  var measurements []int
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    // Convert measurement into integer
    m, err := strconv.Atoi(scanner.Text())
    if err != nil {
      return nil, err
    }
    measurements = append(measurements, m)
  }

  return measurements, scanner.Err()
}

func sum(input []int) int {
  s := 0
  for _, element := range input {
    s = s + element
  }
  return s
}

func part1(measurements []int) int {
  count := 0
  for i := 1; i < len(measurements); i++ {
    if measurements[i - 1] < measurements[i] {
      count = count + 1
    }
  }
  return count
}

func part2(measurements []int) int {
  count := 0
  var _measurements []int

  // Group measurements in tuples
  for i := 0; i < len(measurements) - 2; i++ {
    _measurements = append(_measurements, sum(measurements[i:i+3]))
  }

  // Count increased measurements
  for i := 1; i < len(_measurements); i++ {
    if _measurements[i - 1] < _measurements[i] {
      count = count + 1
    }
  }

  return count
}

func main() {
  measurements, err := readMeasurements("input.txt")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // Part 1
  fmt.Printf("Part 1: Number of increased measurements: %d\n", part1(measurements))

  // Part 2
  fmt.Printf("Part 2: Number of increased measurements: %d\n", part2(measurements))
}

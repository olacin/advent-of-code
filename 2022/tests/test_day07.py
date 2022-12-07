from pathlib import Path

from advent_of_code.day07 import part1, part2

INPUT = Path(__file__).parent / "data" / "day07.txt"


def test_part1():
    # assert part1(INPUT) == 95437
    assert part1(INPUT) == 1886043


def test_part2():
    # assert part2(INPUT) == 24933642
    assert part2(INPUT) == 3842121

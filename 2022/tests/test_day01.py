from pathlib import Path

from advent_of_code.day01 import part1, part2

INPUT = Path(__file__).parent / "data" / "day01.txt"


def test_part1():
    assert part1(INPUT) == 74198


def test_part2():
    assert part2(INPUT) == 209914

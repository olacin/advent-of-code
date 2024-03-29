from pathlib import Path

from advent_of_code.day04 import part1, part2

INPUT = Path(__file__).parent / "data" / "day04.txt"


def test_part1():
    assert part1(INPUT) == 450


def test_part2():
    assert part2(INPUT) == 837

from pathlib import Path

from advent_of_code.day02 import part1, part2

INPUT = Path(__file__).parent / "data" / "day02.txt"


def test_part1():
    assert part1(INPUT) == 13052


def test_part2():
    assert part2(INPUT) == 13693

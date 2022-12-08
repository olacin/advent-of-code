from pathlib import Path

from advent_of_code.day08 import part1, part2

INPUT = Path(__file__).parent / "data" / "day08.txt"


def test_part1():
    # assert part1(INPUT) == 21
    assert part1(INPUT) == 1672


def test_part2():
    assert part2(INPUT) == 327180

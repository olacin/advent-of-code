from pathlib import Path

from advent_of_code.day10 import part1, part2

INPUT = Path(__file__).parent / "data" / "day10.txt"


def test_part1():
    # assert part1(INPUT) == 13140
    assert part1(INPUT) == 14220


def test_part2():
    assert part2(INPUT) == "ZRARLFZU"

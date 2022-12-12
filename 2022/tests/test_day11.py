from pathlib import Path

from advent_of_code.day11 import part1, part2

INPUT = Path(__file__).parent / "data" / "day11.txt"


def test_part1():
    # assert part1(INPUT) == 10605
    assert part1(INPUT) == 61005


def test_part2():
    # assert part2(INPUT) == 2713310158
    assert part2(INPUT) == 20567144694

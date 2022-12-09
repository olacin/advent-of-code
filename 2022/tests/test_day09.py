from pathlib import Path

from advent_of_code.day09 import part1, part2

INPUT = Path(__file__).parent / "data" / "day09.txt"


def test_part1():
    # assert part1(INPUT) == 13
    assert part1(INPUT) == 6522


def test_part2():
    # assert part2(INPUT) == 24933642
    assert part2(INPUT) == 2717

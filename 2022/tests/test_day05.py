from pathlib import Path

from advent_of_code.day05 import part1, part2

INPUT = Path(__file__).parent / "data" / "day05.txt"


def test_part1():
    assert part1(INPUT) == "LBLVVTVLP"


def test_part2():
    assert part2(INPUT) == "TPFFBDRJD"

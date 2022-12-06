from pathlib import Path

from advent_of_code.day06 import find_marker, part1, part2

INPUT = Path(__file__).parent / "data" / "day06.txt"


def test_part1():
    assert part1(INPUT) == 1356


def test_part2():
    assert part2(INPUT) == 2564


def test_find_marker():
    assert find_marker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4) == 7
    assert find_marker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14) == 19
    assert find_marker("mjqjpqmgbljspjdztnvjfqwjcgsmlb", 14) == -1

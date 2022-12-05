from pathlib import Path

import pytest

from advent_of_code.day05 import parse_op, part1, part2

INPUT = Path(__file__).parent / "data" / "day05.txt"


def test_part1():
    assert part1(INPUT) == "LBLVVTVLP"


def test_part2():
    assert part2(INPUT) == "TPFFBDRJD"


def test_failing_parse_op():
    with pytest.raises(ValueError):
        parse_op("another one bites the dust")

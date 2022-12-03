from pathlib import Path

import pytest

from advent_of_code.day02 import choice_score, find_choice, part1, part2, round_score

INPUT = Path(__file__).parent / "data" / "day02.txt"


def test_part1():
    assert part1(INPUT) == 13052


def test_part2():
    assert part2(INPUT) == 13693


def test_misc():
    v1, v2 = "I", "J"
    assert choice_score(v1) == 0
    assert round_score(v1, v2) == 0
    with pytest.raises(ValueError):
        find_choice(v1, v2)

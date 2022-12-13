from pathlib import Path

import pytest

from advent_of_code.day13 import Packet, part1, part2

INPUT = Path(__file__).parent / "data" / "day13.txt"


def test_part1():
    # assert part1(INPUT) == 13
    assert part1(INPUT) == 5366


def test_part2():
    # assert part2(INPUT) == 140
    assert part2(INPUT) == 23391


def test_gt():
    assert Packet([[6]]) > Packet([[2]])
    with pytest.raises(ValueError):
        _ = Packet(dict()) > Packet(dict())

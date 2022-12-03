import string
from pathlib import Path
from typing import List, Tuple

alphabet = string.ascii_lowercase + string.ascii_uppercase


def read_input(input: Path):
    with input.open() as fd:
        for line in fd.readlines():
            yield line.strip()


def splitline(line: str) -> Tuple[str, str]:
    half = len(line) // 2
    return line[:half], line[half:]


def chunks(arr: List[str], size: int):
    for i in range(0, len(arr), size):
        yield arr[i : i + size]  # noqa: E203


def common_item(strings: List[str]) -> str:
    return set.intersection(*map(set, strings)).pop()


def item_score(item: str):
    return alphabet.find(item) + 1


def part1(input: Path) -> int:
    score = 0
    for line in read_input(input):
        c1, c2 = splitline(line)
        item = common_item([c1, c2])
        score += item_score(item)
    return score


def part2(input: Path) -> int:
    score = 0
    lines = list(read_input(input))
    for group in chunks(lines, 3):
        item = common_item(group)
        score += item_score(item)
    return score

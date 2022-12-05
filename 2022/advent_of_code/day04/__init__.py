from pathlib import Path
from typing import List, Tuple


def read_input(input: Path):
    with input.open() as fd:
        for line in fd.readlines():
            yield line.strip()


def to_range(pair: str):
    start, end = pair.split("-")
    return list(range(int(start), int(end) + 1))


def get_sections(line: str) -> Tuple[List[int], List[int]]:
    p1, p2 = line.split(",")
    return to_range(p1), to_range(p2)


def is_contained(l1: List[int], l2: List[int]) -> bool:
    return set(l1).issubset(l2) or set(l2).issubset(l1)


def overlaps(l1: List[int], l2: List[int]):
    return len(set(l1).intersection(l2)) > 0


def part1(input: Path) -> int:
    score = 0
    for line in read_input(input):
        s1, s2 = get_sections(line)
        if is_contained(s1, s2):
            score += 1
    return score


def part2(input: Path) -> int:
    score = 0
    for line in read_input(input):
        s1, s2 = get_sections(line)
        if overlaps(s1, s2):
            score += 1
    return score

from pathlib import Path


def read_input(input: Path):
    with input.open() as fd:
        return fd.readline()


def find_marker(s: str, n: int) -> int:
    high = len(s) - n - 1
    for i in range(high):
        val = s[i : i + n]  # noqa: E203
        if len(set(val)) == n:
            return i + n
    return -1


def part1(input: Path):
    s = read_input(input)
    return find_marker(s, 4)


def part2(input: Path):
    s = read_input(input)
    return find_marker(s, 14)

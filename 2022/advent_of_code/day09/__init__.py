import operator
from pathlib import Path
from typing import List, Tuple

directions = {
    "U": (0, 1),
    "D": (0, -1),
    "R": (1, 0),
    "L": (-1, 0),
}


def read_input(input: Path):
    with input.open() as fd:
        for line in fd.readlines():
            dir, count = line.strip().split()
            yield (dir, int(count))


def move_head(head, symbol):
    return tuple(map(operator.add, head, directions[symbol]))


def move_tail(head, tail) -> Tuple[int, int]:
    hx, hy = head
    tx, ty = tail
    # no need to move the tail
    if abs(tx - hx) <= 1 and abs(ty - hy) <= 1:
        pass
    # horizontal move
    elif tx == hx:
        ty += 1 if hy > ty else -1
    # vertical move
    elif ty == hy:
        tx += 1 if hx > tx else -1
    # diagonal move
    else:
        tx += 1 if hx > tx else -1
        ty += 1 if hy > ty else -1

    return (tx, ty)


def move(rope, direction):
    rope[0] = move_head(rope[0], direction)
    for idx in range(1, len(rope)):
        t = move_tail(rope[idx - 1], rope[idx])
        rope[idx] = t


def find_visited(input: Path, rope_length: int):
    content = read_input(input)
    rope: List[Tuple[int, int]] = [(0, 0) for _ in range(rope_length)]
    positions: set[Tuple[int, int]] = {(0, 0)}
    for dir, count in content:
        for _ in range(count):
            move(rope, dir)
            positions.add(rope[-1])
    return positions


def part1(input: Path):
    return len(find_visited(input, 2))


def part2(input: Path):
    return len(find_visited(input, 10))

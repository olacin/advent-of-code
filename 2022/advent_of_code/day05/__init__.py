import re
from collections import deque
from pathlib import Path
from typing import Deque, List, Tuple


def read_input(input: Path):
    with input.open() as fd:
        data = fd.read()
    p1, p2 = data.split("\n\n")
    stacks = parse_stacks(p1)
    ops = [parse_op(line.strip()) for line in p2.splitlines()]
    return stacks, ops


def parse_stacks(content: str) -> List[Deque[str]]:
    lines = list(reversed(content.splitlines()))
    positions = [m.start(0) for m in re.finditer(r"\d", lines[0])]
    stacks = [deque() for _ in positions]
    for line in lines[1:]:
        for i, idx in enumerate(positions):
            value = line[idx]
            if not value or value == " ":
                continue
            stacks[i].append(value)
    return stacks


def parse_op(op: str) -> Tuple[int, int, int]:
    pattern = r"move (\d+) from (\d) to (\d)"
    result = re.search(pattern, op)
    if not result:
        raise ValueError(f"Unable to parse op: {op}")
    return tuple(map(int, result.groups()))


def top_values(stacks: List[Deque[str]]):
    values = [stack[-1] for stack in stacks if len(stack) > 0]
    return "".join(values)


def exec_ops(input: Path, reverse: bool = False):
    stacks, ops = read_input(input)
    for count, src, dst in ops:
        values = [stacks[src - 1].pop() for _ in range(count)]
        if reverse:
            values = reversed(values)
        stacks[dst - 1].extend(values)
    return top_values(stacks)


def part1(input: Path) -> str:
    return exec_ops(input)


def part2(input: Path) -> str:
    return exec_ops(input, True)

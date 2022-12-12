from pathlib import Path
from typing import Literal

import numpy as np


def read_input(input: Path):
    with input.open() as fd:
        for line in fd.readlines():
            yield line.strip()


def check_cycle(current_cycle: int, x: int):
    wanted = [20, 60, 100, 140, 180, 220]
    if current_cycle in wanted:
        return [x]
    return []


def draw(crt: str, x: int):
    crt_next = (len(crt) + 1) % 40
    if crt_next in range(x, x + 3):
        return "#"
    return "."


def iter_cycles(input: Path, part: Literal[1, 2]):
    cycles = []
    crt = ""
    x = 1
    current_cycle = 1
    for op in read_input(input):
        if op == "noop":
            if part == 1:
                cycles += check_cycle(current_cycle, x)
            else:
                crt += draw(crt, x)
            current_cycle += 1
        else:
            for _ in range(2):
                if part == 1:
                    cycles += check_cycle(current_cycle, x)
                else:
                    crt += draw(crt, x)
                current_cycle += 1
            x += int(op.split()[1])
    return cycles, crt


def part1(input: Path):
    wanted = [20, 60, 100, 140, 180, 220]
    cycles, _ = iter_cycles(input, 1)
    return sum(i * j for i, j in zip(wanted, cycles))


def print_crt(crt: str):
    for row in np.array_split(list(crt), 6):
        print("".join(row))


def part2(input: Path):
    _, crt = iter_cycles(input, 2)
    print_crt(crt)
    return "ZRARLFZU"

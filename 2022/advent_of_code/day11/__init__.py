import math
import re
from operator import add, mul
from pathlib import Path
from typing import Dict, List, Literal

pattern = re.compile(
    r"""Monkey (\d):
  Starting items: (.*)
  Operation: new = old (.) (.*)
  Test: divisible by (\d*)
    If true: throw to monkey (\d)
    If false: throw to monkey (\d)"""
)


class Monkey:
    def __init__(
        self, id: int, items: List[int], op, test: int, cond_true: int, cond_false: int
    ):
        self.id = id
        self.items = items
        self.op = op
        self.test_val = test
        self.cond_true = cond_true
        self.cond_false = cond_false
        self._inspect_calls = 0

    def inspect(self, part: Literal[1, 2], lcm: int = 1):
        operations = []
        if not self.items:
            return operations
        for _ in range(len(self.items)):
            self._inspect_calls += 1
            current = self.items.pop(0)
            new = self.op(current)
            if part == 1:
                new = new // 3
            else:
                new = new % lcm
            dst = self.cond_true if self.test(new) else self.cond_false
            operations.append((dst, new))
        return operations

    def test(self, value: int) -> bool:
        return value % self.test_val == 0

    @staticmethod
    def parse_op(op: str, value: str):
        fn = add if op == "+" else mul
        try:
            val = int(value)
            return lambda old: fn(old, val)
        except ValueError:
            return lambda old: fn(old, old)

    @classmethod
    def from_str(cls, s: str):
        m = pattern.match(s)
        if not m:
            raise ValueError("Unable to parse monkey input")  # no cov
        _id = int(m.group(1))
        _items = [int(x) for x in m.group(2).split(", ")]
        _op = Monkey.parse_op(m.group(3), m.group(4))
        _test = int(m.group(5))
        _cond_true = int(m.group(6))
        _cond_false = int(m.group(7))
        return cls(_id, _items, _op, _test, _cond_true, _cond_false)


def read_input(input: Path):
    with input.open() as fd:
        return fd.read().split("\n\n")


def play_rounds(input: Path, n: int, **kwargs):
    monkeys: Dict[int, Monkey] = {}
    for part in read_input(input):
        monkey = Monkey.from_str(part)
        monkeys[monkey.id] = monkey

    if kwargs.get("part", 1) == 2:
        kwargs["lcm"] = math.lcm(*[m.test_val for m in monkeys.values()])

    for _ in range(n):
        for m in monkeys.values():
            for next_id, value in m.inspect(**kwargs):
                monkeys[next_id].items.append(value)

    x, y = sorted([m._inspect_calls for m in monkeys.values()])[-2:]
    return x * y


def part1(input: Path):
    return play_rounds(input, 20, part=1)


def part2(input: Path):
    return play_rounds(input, 10000, part=2)

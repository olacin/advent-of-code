import functools
import json
from pathlib import Path
from typing import List, Union


def read_input(input: Path):
    with input.open() as fd:
        for pt in fd.read().split("\n\n"):
            yield tuple([Packet(json.loads(line.strip())) for line in pt.splitlines()])


@functools.total_ordering
class Packet:
    def __init__(self, value):
        self.value: Union[int, List] = value

    def __eq__(self, b: "Packet"):
        return self.value == b.value

    def __gt__(self, b: "Packet"):
        # 2 integers
        if isinstance(self.value, int) and isinstance(b.value, int):
            return self.value > b.value
        # integer & list
        if isinstance(self.value, int) and isinstance(b.value, list):
            return Packet([self.value]) > b
        if isinstance(b.value, int) and isinstance(self.value, list):
            return self > Packet([b.value])
        # 2 lists
        if isinstance(self.value, list) and isinstance(b.value, list):
            for v1, v2 in zip(self.value, b.value):
                if v1 != v2:
                    return Packet(v1) > Packet(v2)

            return len(self.value) > len(b.value)

        raise ValueError(
            f"Unsupported types: self={type(self.value)} other={type(b.value)}"
        )

    def __str__(self):
        return str(self.value)  # no cov

    def __repr__(self):
        return str(self)  # no cov


def part1(input: Path):
    pairs = read_input(input)
    order = [idx for idx, (left, right) in enumerate(pairs, 1) if left < right]
    return sum(order)


def part2(input: Path):
    # Flatten pairs into a list of packets
    packets = [packet for pair in read_input(input) for packet in pair]
    dividers = [Packet([[2]]), Packet([[6]])]
    packets.extend(dividers)
    sorted_packets = sorted(packets)
    x = sorted_packets.index(dividers[0])
    y = sorted_packets.index(dividers[1])
    return (x + 1) * (y + 1)

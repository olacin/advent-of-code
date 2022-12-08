from pathlib import Path

import numpy as np


def read_input(input: Path):
    data = []
    with input.open() as fd:
        for line in fd.readlines():
            data.append(list(line.strip()))
    return np.array(data, dtype=str)


def is_visible(arr, x, y) -> bool:
    width, height = arr.shape
    value = arr[x, y]
    # edge
    if (x == 0 or x == width - 1) or (y == 0 or y == height - 1):
        return True
    # interior
    left = arr[x, :y]
    top = arr[:x, y]
    right = arr[x, y+1:]
    bottom = arr[x+1:, y]

    for direction in (left, top, right, bottom):
        if all(elem < value for elem in direction):
            return True
    return False


def count_visible(trees, value):
    count = 0
    for tree in trees:
        count += 1
        if tree >= value:
            break
    return count


def count_visible_around(arr, x, y) -> int:
    width, height = arr.shape
    # edge
    if (x == 0 or x == width - 1) or (y == 0 or y == height - 1):
        return 0
    value = arr[x, y]
    left = count_visible(reversed(arr[x, :y]), value)
    top = count_visible(reversed(arr[:x, y]), value)
    right = count_visible(arr[x, y + 1 :], value)
    bottom = count_visible(arr[x + 1 :, y], value)

    return left * top * right * bottom


def part1(input: Path):
    arr = read_input(input)
    return sum((1 if is_visible(arr, i, j) else 0 for i, j in np.ndindex(arr.shape)))


def part2(input: Path):
    arr = read_input(input)
    return max((count_visible_around(arr, i, j) for i, j in np.ndindex(arr.shape)))

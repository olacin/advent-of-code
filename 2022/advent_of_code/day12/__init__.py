import itertools
import operator
from collections import deque
from pathlib import Path
from typing import Dict, Tuple

import numpy as np
from numpy.typing import ArrayLike

Point = Tuple[int, int]


def valid_neighbors(grid: ArrayLike, visited: Dict[Point, Point], cell):
    def add_tuples(t1: Point, t2: Point):
        return tuple(map(operator.add, t1, t2))

    directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]
    neighbors = [add_tuples(cell, dir) for dir in directions]
    return filter(lambda nb: is_valid(grid, visited, cell, nb), neighbors)


def elevation(value: str):
    match value:
        case "S":
            return ord("a")
        case "E":
            return ord("z")
        case _:
            return ord(value)


def read_input(input: Path):
    matrix = []
    with input.open() as fd:
        for line in fd.readlines():
            matrix.append(list(line.strip()))
    return np.array(matrix)


def find_positions(matrix: ArrayLike, item: str):
    for coordinates in np.argwhere(matrix == item):
        yield tuple(coordinates)


def is_valid(grid, visited: Dict[Point, Point], current: Point, next: Point):
    width, height = grid.shape
    cx, cy = current
    x, y = next
    # Out of bounds
    if x < 0 or y < 0 or x >= width or y >= height:
        return False
    # Visited
    if next in visited:
        return False
    # Height difference
    if elevation(grid[x][y]) - elevation(grid[cx][cy]) > 1:
        return False

    return True


def bfs(grid: ArrayLike, start: Point, end: Point):
    q = deque()
    q.append(start)
    visited = {}
    visited[start] = None
    cell = None

    while q:
        cell = q.popleft()

        if cell == end:
            break

        for pos in valid_neighbors(grid, visited, cell):
            q.append(pos)
            visited[pos] = cell

    # Exit early if there's no valid path
    if cell != end:
        return None

    # Rebuild path
    path = []
    curr = end
    while curr != start:
        path.append(curr)
        curr = visited[curr]

    return path


def part1(input: Path):
    matrix = read_input(input)
    start = next(find_positions(matrix, "S"))
    end = next(find_positions(matrix, "E"))
    if path := bfs(matrix, start, end):
        return len(path)
    return -1  # no cov


def part2(input: Path):
    matrix = read_input(input)
    starts = itertools.chain(find_positions(matrix, "S"), find_positions(matrix, "a"))
    end = next(find_positions(matrix, "E"))
    paths = filter(None, [bfs(matrix, start, end) for start in starts])
    return min(map(len, paths))

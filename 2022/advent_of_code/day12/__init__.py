import operator
from collections import deque
from operator import itemgetter
from pathlib import Path
from typing import Tuple

import numpy as np

DIRECTIONS = [
    (-1, 0),
    (0, 1),
    (1, 0),
    (0, -1),
]


def read_input(input: Path):
    matrix = []
    with input.open() as fd:
        for line in fd.readlines():
            matrix.append(list(line.strip()))
    return np.array(matrix)


def find_position(matrix, item):
    return tuple([itemgetter(0)(coord) for coord in np.where(matrix == item)])


def add_tuples(t1: Tuple, t2: Tuple):
    return tuple(map(operator.add, t1, t2))


def is_valid(grid, visited, current, next):
    width, height = visited.shape
    cx, cy = current
    x, y = next
    # Out of bounds
    if x < 0 or y < 0 or x >= width or y >= height:
        print(f"{(x, y)} is out of bounds")
        return False
    # Visited
    if visited[x][y] is True:
        print(f"{(x, y)} has already been visited")
        return False
    # Height difference
    if grid[cx][cy] != "S" and ord(grid[x][y]) - ord(grid[cx][cy]) > 1:
        print(f"{grid[x][y]} - {grid[cx][cy]} diff is too big")
        return False

    return True


def bfs(grid, visited, start):
    q = deque()
    x, y = start
    q.append(start)
    visited[x][y] = True

    print("I am in BFS")

    while len(q) > 0:
        print(len(q))
        cell = q.popleft()
        x, y = cell
        print(grid[x][y], end=" ")

        if grid[x][y] == "E":
            print("Element has been found")
            return q

        for dir in DIRECTIONS:
            pos = add_tuples(cell, dir)
            print(f"Checking position: {pos}")
            if is_valid(grid, visited, cell, pos):
                print(f"position is valid: {pos}")
                q.append(pos)
                visited[pos[0]][pos[1]] = True
        print(visited)
        # raise ZeroDivisionError()


def part1(input: Path):
    matrix = read_input(input)
    visited = matrix.copy()
    visited.fill(False)
    start = find_position(matrix, "S")
    end = find_position(matrix, "E")
    print("Matrix")
    print(matrix)
    print("Visited")
    print(visited)
    print(f"Start: {start}")
    print(f"End: {end}")
    q = bfs(matrix, visited, start)
    # print(q)
    # raise NotImplementedError()


def part2(input: Path):
    raise NotImplementedError()


if __name__ == "__main__":
    import sys

    part1(Path(sys.argv[1]))

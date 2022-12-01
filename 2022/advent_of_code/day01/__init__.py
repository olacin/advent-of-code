from pathlib import Path


def get_elves_total_calories(input: Path):
    with input.open() as fd:
        data = fd.read()
    for gp in data.split("\n\n"):
        yield sum([int(line) for line in gp.splitlines()])


def part1(input: Path) -> int:
    calories = get_elves_total_calories(input)
    return max(calories)


def part2(input: Path) -> int:
    calories = get_elves_total_calories(input)
    return sum(sorted(calories)[-3:])

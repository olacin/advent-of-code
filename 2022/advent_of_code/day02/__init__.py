from pathlib import Path


def read_input(input: Path):
    with input.open() as fd:
        for line in fd.readlines():
            yield tuple(line.split())


def choice_score(letter: str):
    match letter:
        case "X":
            return 1
        case "Y":
            return 2
        case "Z":
            return 3
        case _:
            return 0


def find_choice(opponent: str, result: str):
    match (opponent, result):
        # Should play rock
        case ("B", "X") | ("A", "Y") | ("C", "Z"):
            return "X"
        # Should play paper
        case ("C", "X") | ("B", "Y") | ("A", "Z"):
            return "Y"
        # Should play scissors
        case ("A", "X") | ("C", "Y") | ("B", "Z"):
            return "Z"
        case _:
            raise ValueError("Unable to make a valid choice")


def round_score(opponent: str, user: str):
    match (opponent, user):
        # Draws
        case ("A", "X") | ("B", "Y") | ("C", "Z"):
            return 3
        # Losses
        case ("A", "Z") | ("B", "X") | ("C", "Y"):
            return 0
        # Wins
        case ("A", "Y") | ("B", "Z") | ("C", "X"):
            return 6
        case _:
            return 0


def part1(input: Path) -> int:
    score = 0
    for opponent, user in read_input(input):
        score += round_score(opponent, user)
        score += choice_score(user)
    return score


def part2(input: Path) -> int:
    score = 0
    for opponent, result in read_input(input):
        choice = find_choice(opponent, result)
        score += round_score(opponent, choice)
        score += choice_score(choice)
    return score

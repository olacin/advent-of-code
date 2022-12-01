from typing import Optional

import httpx
from typer import FileBinaryWrite


def fetch_input(day: int, destination: FileBinaryWrite, cookie: Optional[str]):
    if day < 1 or day > 25:
        raise ValueError("Not a valid day")
    if not cookie:
        raise ValueError("Empty cookie")
    res = httpx.get(
        f"https://adventofcode.com/2022/day/{day}/input", cookies={"session": cookie}
    )
    res.raise_for_status()
    destination.write(res.content)

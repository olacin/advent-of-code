from typing import Optional

import httpx
import typer

app = typer.Typer()


@app.command()
def main(
    day: int,
    output: typer.FileBinaryWrite = typer.Option("-"),
    cookie: Optional[str] = typer.Option(None, envvar="COOKIE"),
):
    if day < 1 or day > 25:
        raise ValueError("Not a valid day")
    if not cookie:
        raise ValueError("Empty cookie")
    url = f"https://adventofcode.com/2022/day/{day}/input"
    res = httpx.get(url, cookies={"session": cookie})
    res.raise_for_status()
    output.write(res.content)


if __name__ == "__main__":
    typer.run(main)

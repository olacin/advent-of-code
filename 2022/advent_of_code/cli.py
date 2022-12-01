from typing import Optional

import typer

from .utils import fetch_input

app = typer.Typer()


@app.command()
def main(
    day: int,
    output: typer.FileBinaryWrite = typer.Option("-"),
    cookie: Optional[str] = typer.Option(None, envvar="COOKIE"),
):
    fetch_input(day, output, cookie)


if __name__ == "__main__":
    typer.run(main)

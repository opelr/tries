"""
"""

from dataclasses import dataclass
from pathlib import Path
from typing import List

import click


@dataclass
class Node:
    value: str
    parent: "Node"
    children: List["Node"]


def read_file_as_list(path: Path) -> List[str]:
    with open(path) as f:
        return [i.split() for i in f.readlines().strip()]


@click.command()
@click.option('--num-words', default=1, help='')
@click.option('--ignore-case', default=True, help='')
@click.argument('file')
def main(num_words, ignore_case, file) -> None:
    pass


if __name__ == "__main__":
    main()

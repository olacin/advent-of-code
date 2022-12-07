from dataclasses import dataclass, field
from pathlib import Path
from typing import Dict, List, Optional


@dataclass
class INode:
    name: str
    kind: str = "dir"
    size: int = 0
    parent: Optional["INode"] = None
    children: Dict[str, "INode"] = field(default_factory=dict)


def read_input(input: Path):
    with input.open() as fd:
        for line in fd.readlines():
            yield line.strip()


def dir_size(node: INode) -> int:
    if node.kind == "file":
        return node.size
    size = 0
    for child in node.children.values():
        size += dir_size(child)
    return size


def build_tree(input: List[str]):
    root = INode("/")
    current_node = root
    # Part 1: Build tree
    for line in input:
        parts = line.split()
        if parts[1] == "ls":
            continue
        elif parts[1] == "cd":
            if parts[2] == "/":
                current_node = root
            elif parts[2] == "..":
                current_node = current_node.parent
            elif parts[2] not in current_node.children.keys():
                current_node = INode(parts[2])  # no cov
            else:
                current_node = current_node.children[parts[2]]
        elif parts[0] == "dir" and parts[1] not in current_node.children:
            current_node.children[parts[1]] = INode(parts[1], parent=current_node)
        else:
            sz = int(parts[0])
            current_node.children[parts[1]] = INode(
                parts[1], kind="file", size=sz, parent=current_node
            )
    # Part 2: Apply sizes to directories
    apply_dir_sizes(root)
    return root


def apply_dir_sizes(node: INode):
    node.size = dir_size(node)
    for child in node.children.values():
        apply_dir_sizes(child)


def get_sizes(node: INode, sizes: List[int] = []) -> List[int]:
    if node.kind == "dir":
        sizes.append(node.size)
    for child in node.children.values():
        get_sizes(child, sizes)
    return sizes


def part1(input: Path):
    content = list(read_input(input))
    tree = build_tree(content)
    sizes = filter(lambda s: s < 100000, get_sizes(tree))
    return sum(sizes)


def part2(input: Path):
    content = list(read_input(input))
    tree = build_tree(content)
    free_space = 70_000_000 - tree.size
    min_deletion = 30_000_000 - free_space
    return min(filter(lambda size: size >= min_deletion, get_sizes(tree)))

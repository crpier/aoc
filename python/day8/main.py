import math
from itertools import cycle
from io import TextIOWrapper
from typing import Literal, TypedDict, cast

Direction = Literal["L", "R"]


class Destination(TypedDict):
    L: str
    R: str


class Map:
    def __init__(self, starter_node: str) -> None:
        self._nodes: dict[str, Destination] = {}
        self._starter: str = starter_node
        self._current_node = starter_node

    def add_node(self, line: str):
        start, destinations = [part.strip() for part in line.split("=")]
        start = start
        destinations = destinations.replace("(", "").replace(")", "")
        destinations = [dest.strip() for dest in destinations.split(",")]
        destination: Destination = {"L": destinations[0], "R": destinations[1]}
        self._nodes[start] = destination

    def apply_direction(self, direction: Direction):
        new_current = self._nodes[self._current_node][direction]
        self._current_node = new_current
        return new_current

    def __repr__(self) -> str:
        return f"Map {self._starter}"


def parse_input(f: TextIOWrapper) -> tuple[list[Direction], Map]:
    directions = f.readline().strip()
    f.readline()
    map = Map("AAA")
    for line in f:
        map.add_node(line)
    return cast(list[Direction], directions), map


def parse_input_multiple_maps(f: TextIOWrapper) -> tuple[list[Direction], list[Map]]:
    directions = f.readline().strip()
    f.readline()
    map_lines: list[str] = []
    starter_nodes: list[str] = []
    for line in f:
        map_lines.append(line)
        if line[2] == "A":
            starter_nodes.append(line[:3])

    maps: list[Map] = []
    for starter in starter_nodes:
        new_map = Map(starter)
        for line in map_lines:
            new_map.add_node(line)
        maps.append(new_map)

    return cast(list[Direction], directions), maps


# with open("example.txt") as f:
# with open("example_2.txt") as f:
# with open("example_3.txt") as f:
with open("input.txt") as f:
    directions, maps = parse_input_multiple_maps(f)

current_node = 0
steps_to_destinations = []

for map_idx, map in enumerate(maps):
    for steps, direction in enumerate(cycle(directions), start=1):
        new_location = map.apply_direction(direction)
        if new_location.endswith("Z"):
            steps_to_destinations.append(steps)
            break

print(math.lcm(*steps_to_destinations))

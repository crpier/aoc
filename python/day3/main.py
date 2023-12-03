from contextlib import suppress
from dataclasses import dataclass
from typing import Generator, TypeAlias


Schematic: TypeAlias = list[list[str | int]]

SCHEMATIC: Schematic = []


@dataclass
class Point:
    x: int
    y: int


def is_symbol(char: str | int):
    # symbol as in, the game considers it a symbol
    if isinstance(char, int):
        return False
    if char == ".":
        return False
    return not char.isalpha()


def parse_schematic(data: str) -> Schematic:
    schematic: Schematic = []
    for line in data.strip().split("\n"):
        new_line = []
        for char in line:
            try:
                new_line.append(int(char))
            except ValueError:
                new_line.append(char)
        schematic.append(new_line)

    return schematic


def print_schematic():
    for line in SCHEMATIC:
        print("".join([str(element) for element in line]))


def get_number_at_location(loc: Point) -> int:
    if not isinstance(SCHEMATIC[loc.x][loc.y], int):
        raise ValueError("Not a number at given location")
    numbers: list[int] = [int(SCHEMATIC[loc.x][loc.y])]
    # go left
    moving_point = Point(loc.x, loc.y)
    while True:
        if moving_point.y <= 0:
            break
        moving_point.y -= 1
        try:
            numbers.insert(0, int(SCHEMATIC[moving_point.x][moving_point.y]))
        except ValueError:
            break

    # go right
    moving_point = Point(loc.x, loc.y)
    while True:
        if moving_point.y >= len(SCHEMATIC[0]) - 1:
            break
        moving_point.y += 1
        try:
            numbers.append(int(SCHEMATIC[moving_point.x][moving_point.y]))
        except ValueError:
            break
    return int("".join([str(number) for number in numbers]))


def schematic_iterator() -> Generator[tuple[Point, str | int], None, None]:
    for x_id, row in enumerate(SCHEMATIC):
        for y_id, el in enumerate(row):
            yield Point(x_id, y_id), el


def get_adjacent_locations(loc: Point) -> list[Point]:
    locations: list[Point] = []
    # above
    if loc.x > 0:
        # above left
        if loc.y > 0:
            locations.append(Point(loc.x - 1, loc.y - 1))
        # straight below
        locations.append(Point(loc.x - 1, loc.y))
        # above right
        if loc.y < len(SCHEMATIC[0]) - 1:
            locations.append(Point(loc.x - 1, loc.y + 1))
    # straight right
    if loc.y < len(SCHEMATIC[0]) - 1:
        locations.append(Point(loc.x, loc.y + 1))
    # below
    if loc.x < len(SCHEMATIC) - 1:
        # below left
        if loc.y > 0:
            locations.append(Point(loc.x + 1, loc.y - 1))
        # straight below
        locations.append(Point(loc.x + 1, loc.y))
        # below right
        if loc.y < len(SCHEMATIC[0]) - 1:
            locations.append(Point(loc.x + 1, loc.y + 1))
    # straight left
    if loc.y > 0:
        locations.append(Point(loc.x, loc.y - 1))

    return locations

def get_part_numbers_sum(part_numbers: set[int]) -> int:
    return sum(list(part_numbers))


def get_gear_ratio(el: int | str, part_numbers: set[int]) -> int:
    if el == "*" and len(part_numbers) == 2:
        gear_ratio = part_numbers.pop()
        gear_ratio *= part_numbers.pop()
        return gear_ratio
    return 0

total = 0

with open("input.txt") as f:
    # with open("example.txt") as f:
    SCHEMATIC = parse_schematic(f.read())
    for location, el in schematic_iterator():
        # Here, we asume that a number isn't adjacent to 2 different symbols
        if is_symbol(el):
            part_numbers: set[int] = set()
            for adjacent_location in get_adjacent_locations(location):
                with suppress(ValueError):
                    part_numbers.add(get_number_at_location(adjacent_location))
            # part1 answer
            # total += get_part_numbers_sum(part_numbers)
            # part2 answer
            total += get_gear_ratio(el, part_numbers)
print(total)

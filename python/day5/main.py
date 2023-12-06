import functools
from io import TextIOWrapper
import multiprocessing

from dataclasses import dataclass
from multiprocessing.pool import Pool
from queue import Queue
from typing import Iterable


@dataclass
class MapLine:
    destination_range_start: int
    source_range_start: int
    range_length: int


@dataclass
class Map:
    lines: list[MapLine]


NUMBERS_COUNT = 1_815_746_760


def get_seeds(f: TextIOWrapper) -> list[int]:
    line = f.readline().strip()
    # leave the file descriptor on a non-empty line
    f.readline()
    numbers_line = line[line.find(":") + 2 :]
    numbers = [int(number) for number in numbers_line.split()]
    return numbers


def get_seeds_part_2(seed_ranges: list[int]):
    for i in range(0, len(seed_ranges), 2):
        start, length = seed_ranges[i], seed_ranges[i + 1]
        # TODO: the last batch is less than the step
        for j in range(start, start + length, 10000000):
            yield range(j, j+10000000)


def parse_map(f: TextIOWrapper) -> Map:
    # first line of the section is the section name
    _ = f.readline()
    new_line = f.readline().strip()
    map_lines: list[MapLine] = []
    while new_line != "":
        map_line = MapLine(*[int(part) for part in new_line.split()])
        map_lines.append(map_line)
        new_line = f.readline().strip()
    return Map(map_lines)


def apply_map(source: int, map: Map):
    for map_line in map.lines:
        if (
            map_line.source_range_start
            <= source
            <= map_line.source_range_start + map_line.range_length
        ):
            return (
                source - map_line.source_range_start + map_line.destination_range_start
            )
    return source

def apply_maps(source: int) -> int:
    for map in MAPS:
        source = apply_map(source, map)
    return source


def get_lowest_location_number(seeds: Iterable[int]):
    lowest_location_number = 99999999999999999
    for idx, seed in enumerate(seeds):
        location = apply_maps(seed)
        lowest_location_number = min(lowest_location_number, location)
        if idx % 10000000 == 0:
            print(
                f"-{lowest_location_number}"
            )
    return lowest_location_number

MAPS = []


with open("input.txt") as f:
    # with open("example.txt") as f:
    seeds = get_seeds(f)
    MAPS = [parse_map(f) for _ in range(7)]

    # print(f"Part 1 result: {get_lowest_location_number(seeds, maps)}")

    min_location = 9999999999999
    seeds = get_seeds_part_2(seeds)
    pool = multiprocessing.Pool()

    for res in pool.imap_unordered(get_lowest_location_number, seeds):
        min_location = min(min_location, res)
        print(f"+{min_location}")

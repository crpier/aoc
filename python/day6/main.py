from io import TextIOWrapper
from dataclasses import dataclass


@dataclass
class Race:
    time: int
    best_distance: int

@dataclass
class Option:
    time_held: int
    distance_made: int

def get_races_part_1(f: TextIOWrapper) -> list[Race]:
    time_line = f.readline().strip()
    times = time_line[time_line.find(":") + 1 :].strip()
    times = [int(time.strip()) for time in times.split()]

    distance_line = f.readline().strip()
    distances = distance_line[distance_line.find(":") + 1 :].strip()
    distances = [int(distance.strip()) for distance in distances.split()]

    assert len(times) == len(distances)

    return [Race(time, distance) for (time, distance) in zip(times, distances)]


def get_race_options(race: Race) -> list[Option]:
    options = []
    for i in range(1, race.time):
        time_left = race.time - i
        distance_made = time_left * i
        options.append(Option(i, distance_made))
    return options

possibilities = 1
# with open("example.txt") as f:
with open("input.txt") as f:
    races = get_races_part_1(f)
    for race in races:
        winning_options = 0
        race_options = get_race_options(race)
        for option in race_options:
            if option.distance_made > race.best_distance:
                winning_options += 1
        possibilities *= winning_options

print(possibilities)

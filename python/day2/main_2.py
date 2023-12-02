from typing import TypedDict


class CubesPick(TypedDict):
    red: int
    green: int
    blue: int


def adjust_round_mins(round: str, original_mins: CubesPick) -> bool:
    for item in round.split(","):
        item = item.strip()
        count, color = item.split()
        count = int(count)
        match color:
            case "red":
                original_mins["red"] = max(original_mins["red"], count)
            case "green":
                original_mins["green"] = max(original_mins["green"], count)
            case "blue":
                original_mins["blue"] = max(original_mins["blue"], count)

    return True


def get_game_power(line: str) -> int:
    # game_no = int(line[5 : line.find(":")])
    results = line[line.find(":") + 2 :]
    rounds = [round.strip() for round in results.split(";")]
    cubes_mins: CubesPick = {"red": 0, "green": 0, "blue": 0}
    for round in rounds:
        adjust_round_mins(round, cubes_mins)

    return cubes_mins["red"] * cubes_mins["green"] * cubes_mins["blue"]


sum = 0
with open("input.txt") as f:
# with open("example.txt") as f:
    for line in f:
        sum += get_game_power(line.strip())
print(sum)

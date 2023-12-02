def round_is_possible(round: str) -> bool:
    for item in round.split(","):
        item = item.strip() 
        count, color = item.split()
        count = int(count)
        if color == "red" and count > 12:
            return False
        if color == "green" and count > 13:
            return False
        if color == "blue" and count > 14:
            return False

    return True


def game_is_possible(line: str) -> int:
    game_no = int(line[5 : line.find(":")])
    results = line[line.find(":") + 2 :]
    rounds = [round.strip() for round in results.split(";")]
    game_possible = True
    for round in rounds:
        if not round_is_possible(round):
            game_possible = False
    if game_possible:
        return game_no
    return 0


sum = 0
with open("input.txt") as f:
# with open("example.txt") as f:
    for line in f:
        sum += game_is_possible(line.strip())
print(sum)

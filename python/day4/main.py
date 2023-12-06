from functools import partial
from typing import DefaultDict


class Card:
    def __init__(self, line: str) -> None:
        self.score = 0
        self.wins = 0

        card_part, numbers_part = line.split(":", maxsplit=1)
        self.card_no = int(card_part[card_part.find(" ") :])
        winning_numbers_string, had_numbers_string = numbers_part.strip().split("|")
        self.winning_numbers = [
            int(number.strip())
            for number in winning_numbers_string.split(" ")
            if number != ""
        ]
        self.had_numbers = [
            int(number.strip())
            for number in had_numbers_string.split(" ")
            if number != ""
        ]
        self.calculate_wins()

    def calculate_wins(self):
        score = 0
        wins = 0
        for number in self.had_numbers:
            if number in self.winning_numbers:
                score = score * 2 if score != 0 else 1
                wins += 1
        self.score = score
        self.wins = wins

    def __repr__(self) -> str:
        return f"Card {self.card_no}"


CARDS: DefaultDict[int, int] = DefaultDict(partial(int))
score_sum = 0

# with open("example.txt") as f:
with open("input.txt") as f:
    for line in f:
        line = line.strip()
        card = Card(line)
        CARDS[card.card_no] += 1
        score_sum += card.score

        for i in range(CARDS[card.card_no]):
            for j in range(card.wins):
                CARDS[card.card_no + j + 1] += 1


print(f"For part 1: {score_sum}")
print(f"For part 2: {sum(CARDS.values())}")

from collections import Counter
from enum import IntEnum
from typing import Self

CARD_STRENGTHS = {
    "A": 12,
    "K": 11,
    "Q": 10,
    "J": 9,
    "T": 8,
    "9": 7,
    "8": 6,
    "7": 5,
    "6": 4,
    "5": 3,
    "4": 2,
    "3": 1,
    "2": 0,
}

CARD_STRENGTHS_WITH_WILDCARDS = {
    "A": 12,
    "K": 11,
    "Q": 10,
    "T": 9,
    "9": 8,
    "8": 7,
    "7": 6,
    "6": 5,
    "5": 4,
    "4": 3,
    "3": 2,
    "2": 1,
    "J": 0,
}


class HAND_TYPE(IntEnum):
    FIVE_OF_A_KIND = 6
    FOUR_OF_A_KIND = 5
    FULL_HOUSE = 4
    THREE_OF_A_KIND = 3
    TWO_PAIR = 2
    ONE_PAIR = 1
    HIGH_CARD = 0


class Hand:
    cards: str
    bid: int
    hand_type: HAND_TYPE
    with_wildcards: bool

    def __init__(self, cards: str, bid: str, *, with_wildcards: bool = False) -> None:
        self.cards = cards
        self.bid = int(bid)
        self.hand_type = self._get_hand_type(self.cards, with_wildcards=with_wildcards)
        self.with_wildcards = with_wildcards

    @staticmethod
    def _get_hand_type(cards: str, *, with_wildcards: bool):
        if with_wildcards and cards == "JJJJJ":
            return HAND_TYPE.FIVE_OF_A_KIND

        counter = Counter()
        for card in cards:
            if with_wildcards and card == "J":
                continue
            else:
                counter[card] += 1
        if with_wildcards:
            # on a J, simply increment the most common card
            most_common_card = counter.most_common(1)[0][0]
            counter[most_common_card] += cards.count("J")

        match counter.most_common(2):
            # on five of a kind there is only one count
            case [(_, 5)]:
                return HAND_TYPE.FIVE_OF_A_KIND
            case (_, 5), _:
                return HAND_TYPE.FIVE_OF_A_KIND
            case (_, 4), (_, _):
                return HAND_TYPE.FOUR_OF_A_KIND
            case (_, 3), (_, 2):
                return HAND_TYPE.FULL_HOUSE
            case (_, 3), (_, _):
                return HAND_TYPE.THREE_OF_A_KIND
            case (_, 2), (_, 2):
                return HAND_TYPE.TWO_PAIR
            case (_, 2), (_, 1):
                return HAND_TYPE.ONE_PAIR
            case _:
                return HAND_TYPE.HIGH_CARD

    def __lt__(self, other: Self) -> bool:
        if self.hand_type.value < other.hand_type.value:
            return True
        elif self.hand_type.value > other.hand_type.value:
            return False

        for card_self, card_other in zip(self.cards, other.cards):
            strengths_map = (
                CARD_STRENGTHS_WITH_WILDCARDS if self.with_wildcards else CARD_STRENGTHS
            )
            if strengths_map[card_self] < strengths_map[card_other]:
                return True
            elif strengths_map[card_self] > strengths_map[card_other]:
                return False

        raise ValueError("Completely equal cards!!")

    def __repr__(self) -> str:
        return f"{self.cards} => {self.hand_type.name} {self.bid}"


def get_score(lines: list[str], *, part_2: bool = False) -> int:
    hands: list[Hand] = []
    for line in lines:
        hands.append(Hand(*line.split(), with_wildcards=part_2))

    winnings = 0
    for idx, hand in enumerate(sorted(hands)):
        score = hand.bid * (idx + 1)
        winnings += score
    return winnings


# print(Hand("JTTTT", "12"))


with open("input.txt") as f:
    # with open("example.txt") as f:
    lines = [line.strip() for line in f]

part_1_score = get_score(lines)
print(f"Part 1: {part_1_score}")

part_2_score = get_score(lines, part_2=True)
print(f"Part 2: {part_2_score}")

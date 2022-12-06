package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

const (
	Lose = "X"
	Draw = "Y"
	Win  = "Z"
)

type Round struct {
	Opponent int
	Me       int
}

func (round Round) getScore() int {
	score := round.Me
	if round.Opponent == round.Me {
		score += 3
	} else if round.Opponent == Rock {
		if round.Me == Paper {
			score += 6
		} else {
			score += 0
		}
	} else if round.Opponent == Paper {
		if round.Me == Rock {
			score += 0
		} else {
			score += 6
		}
	} else if round.Opponent == Scissors {
		if round.Me == Rock {
			score += 6
		} else {
			score += 0
		}
	}
	return score
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getHandFromLetter(letter string) int {
	switch letter {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		panic("Invalid string for hand")
	}
}

func calculateMyHand(opponentHand int, outcome string) int {
	if outcome == Draw {
		return opponentHand
	}
	if outcome == Lose {
		switch opponentHand {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}
	if outcome == Win {
		switch opponentHand {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}
	panic("Invalid input")
}

func getRoundFromLine(line string) Round {
	split := strings.Split(line, " ")
	opponentHand := getHandFromLetter(split[0])
	myHand := calculateMyHand(opponentHand, split[1])
	return Round{opponentHand, myHand}
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	score := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		round := getRoundFromLine(line)
		score += round.getScore()
    fmt.Println(line)
    fmt.Println(round)
    fmt.Println(round.getScore())
	}
	fmt.Println("")
	fmt.Println(score)
}

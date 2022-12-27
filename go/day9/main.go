package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

const (
	L = "L"
	U = "U"
	D = "D"
	R = "R"
)

type Location struct {
	row      int
	col      int
	startRow int
	startCol int
}

func parseLine(line string) (string, int) {
	result := strings.Split(line, " ")
	direction := result[0]
	count, err := strconv.Atoi(result[1])
	checkErr(err)
	return direction, count
}

func applyHeadMove(direction string, head *Location) {
	switch direction {
	case L:
		head.col -= 1
	case R:
		head.col += 1
	case U:
		head.row -= 1
	case D:
		head.row += 1
	}
}

func printRope(rope [][]string, head Location, tails []Location) {
	for i := 0; i < len(rope); i++ {
		for j := 0; j < len(rope); j++ {
			printedChar := false
			if head.row == i && head.col == j {
				fmt.Printf("H ")
				printedChar = true
			}
			if printedChar == false {
				for t := 0; t < len(tails); t++ {
					if tails[t].row == i && tails[t].col == j {
						fmt.Printf("%d ", t+1)
						printedChar = true
						break
					}
				}
			}
			if printedChar == false && i == head.startRow && j == head.startCol {
				fmt.Printf("s ")
				printedChar = true
			}
			if printedChar == false {
				fmt.Printf(". ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func applyTailFollow(head *Location, tail *Location) {
	/*
	   no motion when:
	   . . .
	   . H .
	   . . .
	*/
	if math.Abs(float64(head.row-tail.row)) <= 1 &&
		math.Abs(float64(head.col-tail.col)) <= 1 {
		return
	}
	/*
		  1 motion horizontal when:
			   . . . .
			   T . H .
			   . . . .
	*/
	if head.row == tail.row && math.Abs(float64(head.col-tail.col)) == 2 {
		if head.col > tail.col {
			tail.col += 1
		} else {
			tail.col -= 1
		}
		return
	}
	/*
		  1 motion vertical when:
			   . T .
			   . . .
			   . H .
			   . . .
	*/
	if head.col == tail.col && math.Abs(float64(head.row-tail.row)) == 2 {
		if head.row > tail.row {
			tail.row += 1
		} else {
			tail.row -= 1
		}
		return
	}
	/*
		    diagonal 1 motion when:
			   . . H
			   T . .
			   . . .

			   T . .
			   . . .
			   . H .
	*/
	if (math.Abs(float64(head.row-tail.row)) == 1 &&
		math.Abs(float64(head.col-tail.col)) == 2) ||
		(math.Abs(float64(head.row-tail.row)) == 2 &&
			math.Abs(float64(head.col-tail.col)) == 1) {
		if head.row > tail.row {
			tail.row += 1
		} else {
			tail.row -= 1
		}
		if head.col > tail.col {
			tail.col += 1
		} else {
			tail.col -= 1
		}
		return
	}
	/*
		    diagonal 1 motion when:
			   . . H
			   . . .
			   T . .
	*/
	if math.Abs(float64(head.row-tail.row)) == 2 &&
		math.Abs(float64(head.col-tail.col)) == 2 {
		if head.row > tail.row {
			tail.row += 1
		} else {
			tail.row -= 1
		}
		if head.col > tail.col {
			tail.col += 1
		} else {
			tail.col -= 1
		}
		return
	}
}

func countTailVisits(tracker [][]int) int {
	count := 0
	for i := 0; i < len(tracker); i++ {
		for j := 0; j < len(tracker[i]); j++ {
			if tracker[i][j] > 0 {
				count += 1
			}
		}
	}
	return count
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	const ropeSize = 1001
	const startPos = (ropeSize - 1) / 2
	rope := make([][]string, ropeSize)
	for i := 0; i < len(rope); i++ {
		rope[i] = make([]string, ropeSize)
	}
	tailTracker := make([][]int, ropeSize)
	for i := 0; i < len(tailTracker); i++ {
		tailTracker[i] = make([]int, ropeSize)
	}
	head := Location{startPos, startPos, startPos, startPos}
	tails := make([]Location, 9)
	for i := 0; i < 9; i++ {
		tails[i] = Location{startPos, startPos, startPos, startPos}
	}
	tailTracker[tails[8].row][tails[8].col] += 1

	// printRope(rope, head, tails)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		direction, count := parseLine(line)
		for i := 0; i < count; i++ {
			applyHeadMove(direction, &head)
			applyTailFollow(&head, &tails[0])
			for i := 1; i < 9; i++ {
				applyTailFollow(&tails[i-1], &tails[i])
			}
			tailTracker[tails[8].row][tails[8].col] += 1
			// printRope(rope, head, tails)
		}
	}
	fmt.Println(countTailVisits(tailTracker))
}

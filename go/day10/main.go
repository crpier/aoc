package main

import (
	"bufio"
	"fmt"
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

func parseCommand(line string, history []int) []int {
	if line == "noop" {
		history = append(history, history[len(history)-1])
	} else {
		add := strings.Split(line, " ")[1]
		addx, err := strconv.Atoi(add)
		checkErr(err)
		history = append(history, history[len(history)-1])
		history = append(history, history[len(history)-1]+addx)
	}
	return history
}

func getRelevantSignals(history []int) (sum int) {
  sum += history[19]*20
  sum += history[59]*60
  sum += history[99]*100
  sum += history[139]*140
  sum += history[179]*180
  sum += history[219]*220
  return sum
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()
	cycleHistory := make([]int, 0)
	cycleHistory = append(cycleHistory, 1)

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		cycleHistory = parseCommand(line, cycleHistory)
	}
	fmt.Println(getRelevantSignals(cycleHistory))
}

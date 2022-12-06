package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getElfCalories(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return sum
		}
		lineNumber, err := strconv.Atoi(line)
		checkErr(err)
		sum += lineNumber
	}
	return -1
}

func maxesOrder(maxContender int, maxesArr *[3]int) {
	for i := 0; i < 3; i++ {
		if maxContender > maxesArr[i] {
			maxesArr[i] = maxContender
			break
		}
	}
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	calorieMaxes := [3]int{-3, -2, -1}
	for {
		elfCalories := getElfCalories(fileScanner)
		if elfCalories == -1 {
			break
		}
		maxesOrder(elfCalories, &calorieMaxes)
	}

	fmt.Println(calorieMaxes)
	fmt.Println(calorieMaxes[0] + calorieMaxes[1] + calorieMaxes[2])
}

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

type Assignment struct {
	firstStart  int
	firstEnd    int
	secondStart int
	secondEnd   int
}

func assignmentFromString(line string) Assignment {
	ranges := strings.Split(line, ",")
	first := strings.Split(ranges[0], "-")
	second := strings.Split(ranges[1], "-")
	firstStart, err := strconv.Atoi(first[0])
	checkErr(err)
	firstEnd, err := strconv.Atoi(first[1])
	checkErr(err)
	secondStart, err := strconv.Atoi(second[0])
	checkErr(err)
	secondEnd, err := strconv.Atoi(second[1])
	checkErr(err)
	return Assignment{firstStart, firstEnd, secondStart, secondEnd}
}

func (asg Assignment) checkOneRangeCoversTheOther() bool {
	if asg.firstStart >= asg.secondStart && asg.firstEnd <= asg.secondEnd {
		return true
	}
	if asg.secondStart >= asg.firstStart && asg.secondEnd <= asg.firstEnd {
		return true
	}
	return false
}

func rangeContains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func (asg Assignment) checkOverlappingRanges() bool {
  firstRange := make([]int, 0)
  secondRange := make([]int, 0)
  for i:= asg.firstStart; i<=asg.firstEnd; i++ {
    firstRange = append(firstRange, i)
  }
  for i:= asg.secondStart; i<=asg.secondEnd; i++ {
    secondRange = append(secondRange, i)
  }
  for _, item := range firstRange {
    if rangeContains(secondRange, item) {
      return true
    }
  }
	return false
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	contains := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		teamAssignment := assignmentFromString(line)
		if teamAssignment.checkOverlappingRanges() == true {
			contains += 1
      fmt.Println(teamAssignment)
		}
	}
	fmt.Println(contains)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func charsAreUnique(chars string) bool {
	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] == chars[j] {
				return false
			}
		}
	}
	return true
}

func findPacketIndex(line string) int {
	for idx := range line {
		if idx < 4 {
			continue
		}
		if charsAreUnique(line[idx-4 : idx]) {
			return idx
		}
	}
	return 0
}

func findMessageIndex(line string) int {
	for idx := range line {
		if idx < 14 {
			continue
		}
		if charsAreUnique(line[idx-14 : idx]) {
			return idx
		}
	}
	return 0
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(findMessageIndex(line))
	}
}

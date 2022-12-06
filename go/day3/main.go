package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func findCommonItem(firstCompartment, secondCompartment string) rune {
	for _, item := range firstCompartment {
		if strings.ContainsRune(secondCompartment, item) {
			return item
		}
	}
	panic("Bad input!")
}

func getItemPriority(item rune) int {
	// lowercase letter
	if item >= 97 && item <= 122 {
		return int(item - 96)
		// upppercase letter
	} else if item >= 65 && item <= 90 {
		return int(item - 38)
	}
	panic("Bad input!")
}

// returns a rune and if EOF
func getTeamBadge(scanner *bufio.Scanner) (rune, bool) {
	lines := [3]string{"", "", ""}
	for i := 0; i <= 2; i++ {
		ok := scanner.Scan()
		if ok == false {
			return '-', true
		}
		lines[i] = scanner.Text()
	}

	for _, item := range lines[0] {
		if strings.ContainsRune(lines[1], item) && strings.ContainsRune(lines[2], item) {
			return item, false
		}
	}
	return '-', true
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for {
		badge, eof := getTeamBadge(fileScanner)
		if eof == true {
			break
		}
		sum += getItemPriority(badge)
		fmt.Printf("%c - %d\n", badge, getItemPriority(badge))
	}

	fmt.Println(sum)
}

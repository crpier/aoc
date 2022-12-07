package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	count int
	from  int
	to    int
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getStackSize(scanner *bufio.Scanner) int {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line[1] == '1' {
			return (len(line) + 1) / 4
		}
	}
	panic("Invalid input")
}

// When this ends, scanner is at an empty line
func parseArrangement(scanner *bufio.Scanner, stackSize int) [][]byte {
	lineStacks := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == '1' {
			break
		}
		lineStack := make([]byte, 0)
		for i := 0; i < stackSize; i++ {
			itemIdx := i*4 + 1
			lineStack = append(lineStack, line[itemIdx])
		}
		lineStacks = append(lineStacks, lineStack)
	}
	return lineStacks
}

func rotateArrangement(matrix [][]byte) [][]byte {
	newMatrix := make([][]byte, 0)
	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	// transpose it
	// for i := range matrix {
	// 	fmt.Printf("%c\n", matrix[i])
	// }
	for i := 0; i < len(matrix[0]); i++ {
		newMatrix = append(newMatrix, make([]byte, len(matrix)))
		for j := 0; j < len(matrix); j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}
	// fmt.Println()
	// for i := range newMatrix {
	// 	fmt.Printf("%c\n", newMatrix[i])
	// }
	return newMatrix
}

func trimStacks(matrix [][]byte) [][]byte {
	newMatrix := make([][]byte, 0)
	for _, stack := range matrix {
		for stack[len(stack)-1] == ' ' {
			stack = stack[:len(stack)-1]
		}
		newMatrix = append(newMatrix, stack)
	}
	return newMatrix
}

func getCommands(scanner *bufio.Scanner, initialArrangement [][]byte) []Command {
	commands := make([]Command, 0)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " ")
		count, err := strconv.Atoi(res[1])
		checkErr(err)
		input_from, err := strconv.Atoi(res[3])
		checkErr(err)
		input_to, err := strconv.Atoi(res[5])
		checkErr(err)
		command := Command{count: count, from: input_from - 1, to: input_to - 1}
		commands = append(commands, command)
	}
	return commands
}

func applyCommand(stacks [][]byte, command Command) {
  movedCrates := stacks[command.from][len(stacks[command.from])-command.count:]
	stacks[command.from] = stacks[command.from][:len(stacks[command.from])-command.count]
	for _, movedCrate := range movedCrates {
		stacks[command.to] = append(stacks[command.to], movedCrate)
	}
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	stackSize := getStackSize(fileScanner)
	fd.Close()

	fd, err = os.Open("input.txt")
	defer fd.Close()
	checkErr(err)

	fileScanner = bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)
	initialArrangement := parseArrangement(fileScanner, stackSize)
	transposedArrangement := rotateArrangement(initialArrangement)
	stacks := trimStacks(transposedArrangement)
	// for i := range stacks {
	// 	fmt.Printf("%c\n", stacks[i])
	// }
	// scan the empty line
	fileScanner.Scan()
	commands := getCommands(fileScanner, initialArrangement)
	for _, command := range commands {
		applyCommand(stacks, command)
		// fmt.Println()
		//   fmt.Println(command)
		// for i := range stacks {
		// 	fmt.Printf("%c\n", stacks[i])
		// }
	}
	fmt.Println()
	for i := range stacks {
		fmt.Printf("%c\n", stacks[i])
	}
	for i := range stacks {
		fmt.Printf("%c", stacks[i][len(stacks[i])-1])
	}
}

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

func getForest(scanner *bufio.Scanner) [][]int {
	forest := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		newRow := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			treeHeight, err := strconv.Atoi(line[i : i+1])
			checkErr(err)
			newRow[i] = treeHeight
		}
		forest = append(forest, newRow)
	}
	return forest
}

func treeIsVisible(row, col int, forest [][]int) bool {
	forestWidth := len(forest[0])
	forestHeight := len(forest)
	// Check if tree is on edge
	if row == 0 || row == forestHeight-1 || col == 0 || col == forestWidth-1 {
		return true
	}
	// Check the sides
	var isVisibleTop = true
	var isVisibleBot = true
	var isVisibleLeft = true
	var isVisibleRight = true
	// Check to the top
	for i := row - 1; i >= 0; i-- {
		if forest[i][col] >= forest[row][col] {
			isVisibleTop = false
		}
	}
	// Check to the bottom
	for i := row + 1; i < forestHeight; i++ {
		if forest[i][col] >= forest[row][col] {
			isVisibleBot = false
		}
	}
	// Check the left
	for j := col - 1; j >= 0; j-- {
		if forest[row][j] >= forest[row][col] {
			isVisibleLeft = false
		}
	}
	// Check the right
	for j := col + 1; j < forestWidth; j++ {
		if forest[row][j] >= forest[row][col] {
			isVisibleRight = false
		}
	}
	return isVisibleTop || isVisibleBot || isVisibleLeft || isVisibleRight
}

func calculateScenicScore(row, col int, forest [][]int) int {
	forestWidth := len(forest[0])
	forestHeight := len(forest)
	// Check if tree is on edge
	if row == 0 || row == forestHeight-1 || col == 0 || col == forestWidth-1 {
		return 0
	}
	// Check the sides
	var scoreTop = 0
	var scoreBot = 0
	var scoreLeft = 0
	var scoreRight = 0
	// Check to the top
	for i := row - 1; i >= 0; i-- {
		if forest[i][col] >= forest[row][col] {
			scoreTop = row - i
			break
		}
		if i == 0 {
			scoreTop = row
		}
	}
	// Check to the bottom
	for i := row + 1; i < forestHeight; i++ {
		if forest[i][col] >= forest[row][col] {
			scoreBot = i - row
			break
		}
		if i == forestHeight-1 {
			scoreBot = forestHeight - row - 1
		}
	}
	// Check the left
	for j := col - 1; j >= 0; j-- {
		if forest[row][j] >= forest[row][col] {
			scoreLeft = col - j
			break
		}
		if j == 0 {
			scoreLeft = col
		}
	}
	// Check the right
	for j := col + 1; j < forestWidth; j++ {
		if forest[row][j] >= forest[row][col] {
			scoreRight = j - col
			break
		}
		if j == forestWidth-1 {
			scoreRight = forestWidth - col - 1
		}
	}
	// fmt.Println(scoreTop)
	// fmt.Println(scoreBot)
	// fmt.Println(scoreLeft)
	// fmt.Println(scoreRight)
	return scoreTop * scoreBot * scoreLeft * scoreRight
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)
	forest := getForest(fileScanner)

	max := -1
	for i := range forest {
		for j, tree := range forest[i] {
			scenicScore := calculateScenicScore(i, j, forest)
			fmt.Printf("%d at %d:%d has score: %d\n", tree, i, j, scenicScore)
			if scenicScore > max {
				max = scenicScore
			}
		}
	}
	fmt.Println(max)

	// fmt.Printf("%d at %d:%d has score: %d\n", forest[3][2], 3, 2, calculateScenicScore(3, 2, forest))
	// fmt.Printf("%d at %d:%d has score: %d\n", forest[1][2], 1, 2, calculateScenicScore(1, 2, forest))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	File = "File"
	Dir  = "Dir"
	ls   = "ls"
	cd   = "cd"
)

type Node struct {
	name     string
	kind     string
	size     int
	parent   *Node
	children *[]Node
}

func (node *Node) repr() string {
	var details string
	if node.kind == Dir {
		details = Dir
	} else {
		details = fmt.Sprintf("%s, size=%d", File, node.size)
	}
	return fmt.Sprintf("- %s (%s)", node.name, details)
}

func (node *Node) printTree(padding int) {
	spaces := strings.Repeat(" ", padding)
	fmt.Printf("%s%s\n", spaces, node.repr())
	for _, child := range *node.children {
		child.printTree(padding + 2)
	}
}

func (node *Node) getNodeSize() int {
	sum := 0
	if node.kind == File {
		sum += node.size
	} else {
		for _, child := range *node.children {
			sum += child.getNodeSize()
		}
	}
	return sum
}

func (node *Node) getSubDirsWithSizeBelow(threshold int) []Node {
	list := make([]Node, 0)
	if node.kind == File {
		panic("Called function for dir nodes on file node")
	} else {
		dirSize := node.getNodeSize()
		if dirSize <= threshold {
			list = append(list, *node)
		}
		for _, child := range *node.children {
			if child.kind == Dir {
				subList := child.getSubDirsWithSizeBelow(threshold)
				if subList != nil {
					list = append(list, subList...)
				}
			}
		}
		return list
	}
}

func (node *Node) getSubDirsWithSizeAbove(threshold int) []Node {
	list := make([]Node, 0)
	if node.kind == File {
		panic("Called function for dir nodes on file node")
	} else {
		dirSize := node.getNodeSize()
		if dirSize >= threshold {
			list = append(list, *node)
		}
		for _, child := range *node.children {
			if child.kind == Dir {
				subList := child.getSubDirsWithSizeAbove(threshold)
				if subList != nil {
					list = append(list, subList...)
				}
			}
		}
		return list
	}
}

func processlsOutput(scanner *bufio.Scanner, currentNode *Node) string {
	children := make([]Node, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			return line
		}
		res := strings.Split(line, " ")
		var kind string
		var size int
		var err error
		if res[0] == "dir" {
			kind = Dir
			size = 0
		} else {
			kind = File
			size, err = strconv.Atoi(res[0])
			checkErr(err)
		}
		subChildren := make([]Node, 0)
		child := Node{name: res[1], kind: kind, size: size, parent: currentNode, children: &subChildren}
		children = append(children, child)
		currentNode.children = &children
	}
	return ""
}

func getNewCWDFromcd(cdArg string, currentNode *Node) *Node {
	if cdArg == ".." {
		return currentNode.parent
	} else {
		for i := 0; i < len(*currentNode.children); i++ {
			children := *currentNode.children
			if children[i].name == cdArg {
				childAddr := &children[i]
				return childAddr
			}
		}
	}
	panic("Bad input")
}

func main() {
	fd, err := os.Open("input.txt")
	checkErr(err)
	defer fd.Close()

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	rootChildren := make([]Node, 0)
	rootNode := Node{kind: Dir, children: &rootChildren, name: "/"}
	currentNode := &rootNode

	// Getting to root node
	fileScanner.Scan()
	fileScanner.Text()

	// First unknown command
	fileScanner.Scan()
	line := fileScanner.Text()
	nextCommand := strings.Split(line, " ")
	for {
		if nextCommand[1] == ls {
			nextLine := processlsOutput(fileScanner, currentNode)
			if nextLine == "" {
				break
			}
			nextCommand = strings.Split(nextLine, " ")
		} else if nextCommand[1] == cd {
			cdArg := nextCommand[2]
			currentNode = getNewCWDFromcd(cdArg, currentNode)
			fileScanner.Scan()
			nextLine := fileScanner.Text()
			nextCommand = strings.Split(nextLine, " ")
		}
	}
	rootNode.printTree(0)
	minToDelete := 30000000 - (70000000 - rootNode.getNodeSize())
	// res := rootNode.getSubDirsWithSizeBelow(100000)
	res := rootNode.getSubDirsWithSizeAbove(minToDelete)
	// sort by size DESC
	sort.Slice(res, func(i, j int) bool { return res[i].getNodeSize() > res[j].getNodeSize() })
	for _, node := range res {
		fmt.Printf("%s - size %d\n", node.name, node.getNodeSize())
	}
}

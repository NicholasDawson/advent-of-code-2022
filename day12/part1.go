package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("day12/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(file)

	var startingNode, endingNode *Node
	var grid [][]*Node

	var line string
	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		line = scanner.Text()
		var row []*Node
		for colIndex, char := range line {
			newNode := NewNode(CharToHeight(char), rowIndex, colIndex)
			if char == 'S' {
				newNode.shortestPath = 0
				startingNode = newNode
			} else if char == 'E' {
				endingNode = newNode
			}

			row = append(row, newNode)
		}
		grid = append(grid, row)
		rowIndex++
	}
	gridXLen := len(grid)
	gridYLen := len(grid[0])

	var neighborAddPoints = [4]image.Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for toExplore := []*Node{startingNode}; len(toExplore) > 0; {
		node := toExplore[0] // pop
		if node == endingNode {
			break
		}
		toExplore = toExplore[1:] // get rid of first element

		for _, pointToAdd := range neighborAddPoints {
			// If neighbor node is valid point in grid
			neighborNodeX := node.point.X + pointToAdd.X
			neighborNodeY := node.point.Y + pointToAdd.Y
			if neighborNodeX >= 0 && neighborNodeY >= 0 && neighborNodeX < gridXLen && neighborNodeY < gridYLen {
				neighborNode := grid[neighborNodeX][neighborNodeY]
				// If neighbor node can be traveled to
				if neighborNode.height-node.height <= 1 {
					if !neighborNode.visited {
						neighborNode.visited = true
						toExplore = append(toExplore, neighborNode)
					}

					// Now replace neighbor shortestPath with current shortest path if less
					shortestPathToNeighborNodeFromCurrentNode := node.shortestPath + 1
					if shortestPathToNeighborNodeFromCurrentNode < neighborNode.shortestPath {
						neighborNode.shortestPath = shortestPathToNeighborNodeFromCurrentNode
					}
				}
			}
		}
	}

	fmt.Printf("Day 12 - Part 1: %d\n", endingNode.shortestPath)
}

type Node struct {
	shortestPath int
	height       int
	point        image.Point
	visited      bool
}

func CharToHeight(char rune) int {
	if char == 'S' {
		char = 'a'
	} else if char == 'E' {
		char = 'z'
	}
	return int(char - 'a')
}

func NewNode(height, row, col int) *Node {
	return &Node{
		shortestPath: math.MaxInt,
		height:       height,
		point:        image.Point{X: row, Y: col},
		visited:      false,
	}
}

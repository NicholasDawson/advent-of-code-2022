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

	var endingNode image.Point
	var gridHeights [][]int

	var line string
	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		line = scanner.Text()
		var row []int
		for colIndex, char := range line {
			newNode := CharToHeight(char)
			if char == 'E' {
				endingNode = image.Point{X: rowIndex, Y: colIndex}
			}

			row = append(row, newNode)
		}
		gridHeights = append(gridHeights, row)
		rowIndex++
	}

	shortestPath := math.MaxInt
	for x, row := range gridHeights {
		for y, height := range row {
			if height == 0 { // Is 'a'
				newGrid := makeNewGrid(x, y, gridHeights)
				shortestPathForCurrentStartingNode :=
					findShortestPathFromStartingPoint(newGrid[x][y], newGrid[endingNode.X][endingNode.Y], newGrid)
				if shortestPathForCurrentStartingNode < shortestPath {
					shortestPath = shortestPathForCurrentStartingNode
				}
			}
		}
	}

	fmt.Printf("Day 12 - Part 2: %d\n", shortestPath)
}

func makeNewGrid(x, y int, gridHeights [][]int) [][]*Node {
	var grid [][]*Node

	for rowIndex, heightsRow := range gridHeights {
		var row []*Node
		for colIndex, height := range heightsRow {
			newNode := NewNode(height, rowIndex, colIndex)
			if rowIndex == x && colIndex == y {
				newNode.shortestPath = 0
			}

			row = append(row, newNode)
		}
		grid = append(grid, row)
	}
	return grid
}

func findShortestPathFromStartingPoint(startingNode, endingNode *Node, grid [][]*Node) int {
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
	return endingNode.shortestPath
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

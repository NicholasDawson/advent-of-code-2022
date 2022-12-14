package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day08/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var forest [][]int32

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		forestRow := make([]int32, len(line))
		for treeIndex, tree := range line {
			forestRow[treeIndex] = tree - '0'
		}
		forest = append(forest, forestRow)
	}

	fmt.Printf("Day 8 - Part 1: %d\n", howManyTreesVisible(forest))

	highestScenicScore := 0
	for x := range forest {
		for y := range forest[x] {
			currentTreeScore := scenicScore(forest, x, y)
			if currentTreeScore > highestScenicScore {
				highestScenicScore = currentTreeScore
			}
		}
	}
	fmt.Printf("Day 8 - Part 2: %d\n", highestScenicScore)
}

func sliceContains(slice [][2]int, cord [2]int) bool {
	for _, element := range slice {
		if element == cord {
			return true
		}
	}
	return false
}

func howManyTreesVisible(forest [][]int32) int {
	var indexesCounted [][2]int

	var forestColumns = make([][]int32, len(forest[0]))
	for i := range forestColumns {
		forestColumns[i] = make([]int32, len(forest))
	}

	for rowIndex, row := range forest {
		var tallestRowTree int32 = -1
		// left to right
		for colIndex, tree := range row {
			if tree > tallestRowTree {
				tallestRowTree = tree
				indexesCounted = append(indexesCounted, [2]int{rowIndex, colIndex})
			}

			// Build columns array
			forestColumns[colIndex][rowIndex] = tree
		}

		// right to left
		tallestRowTree = -1
		for colIndex := len(row) - 1; colIndex >= 0; colIndex-- {
			if row[colIndex] > tallestRowTree {
				tallestRowTree = row[colIndex]
				if !sliceContains(indexesCounted, [2]int{rowIndex, colIndex}) {
					indexesCounted = append(indexesCounted, [2]int{rowIndex, colIndex})
				}
			}
		}
	}

	for colIndex, col := range forestColumns {
		var tallestColTree int32 = -1
		// top to bottom
		for rowIndex, tree := range col {
			if tree > tallestColTree {
				tallestColTree = tree
				if !sliceContains(indexesCounted, [2]int{rowIndex, colIndex}) {
					indexesCounted = append(indexesCounted, [2]int{rowIndex, colIndex})
				}
			}
		}

		// bottom to top
		tallestColTree = -1
		for rowIndex := len(col) - 1; rowIndex >= 0; rowIndex-- {
			if col[rowIndex] > tallestColTree {
				tallestColTree = col[rowIndex]
				if !sliceContains(indexesCounted, [2]int{rowIndex, colIndex}) {
					indexesCounted = append(indexesCounted, [2]int{rowIndex, colIndex})
				}
			}
		}
	}

	return len(indexesCounted)
}

func scenicScore(forest [][]int32, rowIndex, colIndex int) int {
	// Get row and col of desired tree
	var row = forest[rowIndex]
	var col = make([]int32, len(forest))
	for index := range forest {
		col[index] = forest[index][colIndex]
	}

	// The tree you are getting the scenic score of
	tree := forest[rowIndex][colIndex]

	// to right of tree
	right := 0
	for index := colIndex + 1; index <= len(row)-1; index++ {
		right++
		if row[index] >= tree {
			break
		}
	}

	// to left of tree
	left := 0
	for index := colIndex - 1; index >= 0; index-- {
		left++
		if row[index] >= tree {
			break
		}
	}

	// below tree
	below := 0
	for index := rowIndex + 1; index <= len(col)-1; index++ {
		below++
		if col[index] >= tree {
			break
		}
	}

	// above tree
	above := 0
	for index := rowIndex - 1; index >= 0; index-- {
		above++
		if col[index] >= tree {
			break
		}
	}

	return left * right * below * above
}

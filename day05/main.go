package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("day05/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	blockStacks := make([][]rune, 9)

	var line string
	scanner := bufio.NewScanner(file)

	// Parse blocks
	for scanner.Scan() {
		line = scanner.Text()
		if line[1] == '1' {
			scanner.Scan() // Remove newline before commands
			break
		}

		// Parse line of blocks
		for index, char := range line[1:] {
			if index%4 == 0 && char != ' ' { // Every 4 chars starting on first char
				stackNumber := index / 4
				blockStacks[stackNumber] = append([]rune{char}, blockStacks[stackNumber]...) // prepend
			}
		}
	}

	var moveTimes, fromStack, toStack int
	var blockToMove rune

	// Parse/Execute commands
	for scanner.Scan() {
		line = scanner.Text()

		parsedCommandArray := strings.Split(line, " ")
		moveTimes, _ = strconv.Atoi(parsedCommandArray[1])
		fromStack, _ = strconv.Atoi(parsedCommandArray[3])
		toStack, _ = strconv.Atoi(parsedCommandArray[5])

		// Decrement from and to stacks by 1 to make them same as slice index
		fromStack--
		toStack--

		for x := 0; x < moveTimes; x++ {
			fromStackLen := len(blockStacks[fromStack])
			blockToMove, blockStacks[fromStack] = blockStacks[fromStack][fromStackLen-1], blockStacks[fromStack][:fromStackLen-1] // Pop block from top of from stack
			blockStacks[toStack] = append(blockStacks[toStack], blockToMove)                                                      // Put block on top of to stack
		}
	}

	fmt.Print("Day 5 - Part 1: ")
	for _, stack := range blockStacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}
	fmt.Println()
}

func partTwo() {
	file, err := os.Open("day05/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	blockStacks := make([][]rune, 9)

	var line string
	scanner := bufio.NewScanner(file)

	// Parse blocks
	for scanner.Scan() {
		line = scanner.Text()
		if line[1] == '1' {
			scanner.Scan() // Remove newline before commands
			break
		}

		// Parse line of blocks
		for index, char := range line[1:] {
			if index%4 == 0 && char != ' ' { // Every 4 chars starting on first char
				stackNumber := index / 4
				blockStacks[stackNumber] = append([]rune{char}, blockStacks[stackNumber]...) // prepend
			}
		}
	}

	var moveTimes, fromStack, toStack int
	var blocksToMove []rune

	// Parse/Execute commands
	for scanner.Scan() {
		line = scanner.Text()

		parsedCommandArray := strings.Split(line, " ")
		moveTimes, _ = strconv.Atoi(parsedCommandArray[1])
		fromStack, _ = strconv.Atoi(parsedCommandArray[3])
		toStack, _ = strconv.Atoi(parsedCommandArray[5])

		// Decrement from and to stacks by 1 to make them same as slice index
		fromStack--
		toStack--

		splitIndex := len(blockStacks[fromStack]) - moveTimes
		blocksToMove, blockStacks[fromStack] = blockStacks[fromStack][splitIndex:], blockStacks[fromStack][:splitIndex] // Pop x blocks from top of from stack
		blockStacks[toStack] = append(blockStacks[toStack], blocksToMove...)                                            // Put blocks on top of to stack
	}

	fmt.Print("Day 5 - Part 2: ")
	for _, stack := range blockStacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}
	fmt.Println()
}

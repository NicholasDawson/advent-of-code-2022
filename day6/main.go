package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Day 6 - Part 1: %d\n", partOne())
	fmt.Printf("Day 6 - Part 2: %d\n", partTwo())
}

func sliceContains(slice []uint8, char uint8) bool {
	for _, element := range slice {
		if element == char {
			return true
		}
	}
	return false
}

func partOne() int {
	file, err := os.Open("day6/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	const windowSize = 4

	for index, _ := range line {
		window := make([]uint8, windowSize)
		charsAdded := 0
		for windowIndex := index; windowIndex < (index + windowSize); windowIndex++ {
			currentChar := line[windowIndex]
			if !sliceContains(window, currentChar) {
				window = append(window, currentChar)
				charsAdded++
			} else {
				break
			}
		}
		if charsAdded == windowSize {
			return index + windowSize
		}
	}
	return -1
}

func partTwo() int {
	file, err := os.Open("day6/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	const windowSize = 14

	for index, _ := range line {
		window := make([]uint8, windowSize)
		charsAdded := 0
		for windowIndex := index; windowIndex < (index + windowSize); windowIndex++ {
			currentChar := line[windowIndex]
			if !sliceContains(window, currentChar) {
				window = append(window, currentChar)
				charsAdded++
			} else {
				break
			}
		}
		if charsAdded == windowSize {
			return index + windowSize
		}
	}
	return -1
}

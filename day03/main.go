package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func getLetterPriority(letter rune) int32 {
	if letter >= 'a' {
		return letter - 'a' + 1
	} else {
		return letter - 'A' + 27
	}
}

func partOne() {
	file, err := os.Open("day03/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var line string
	var lineLen int
	scanner := bufio.NewScanner(file)

	var totalScore int32 = 0
	for scanner.Scan() {
		line = scanner.Text()
		lineLen = len(line)

		firstCompartment := line[:lineLen/2]
		secondCompartment := line[lineLen/2:]

		for _, letter := range firstCompartment {
			if strings.ContainsRune(secondCompartment, letter) {
				totalScore += getLetterPriority(letter)
				break
			}
		}
	}

	fmt.Printf("Day 3 - Part 1: %d\n", totalScore)
}

func partTwo() {
	file, err := os.Open("day03/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var line1 string
	var line2 string
	var line3 string
	scanner := bufio.NewScanner(file)

	var totalScore int32 = 0
	for scanner.Scan() {
		line1 = scanner.Text()

		scanner.Scan()
		line2 = scanner.Text()

		scanner.Scan()
		line3 = scanner.Text()

		for _, letter := range line1 {
			if strings.ContainsRune(line2, letter) && strings.ContainsRune(line3, letter) {
				totalScore += getLetterPriority(letter)
				break
			}
		}
	}

	fmt.Printf("Day 3 - Part 2: %d\n", totalScore)
}

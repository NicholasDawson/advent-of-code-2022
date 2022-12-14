package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("day01/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	mostCalories := 0
	currentElfCalories := 0
	var currentLine string
	var currentInt int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine = scanner.Text()
		if currentLine == "" {
			if currentElfCalories > mostCalories {
				mostCalories = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			currentInt, err = strconv.Atoi(currentLine)
			if err != nil {
				log.Fatalf("Bad string to int conversion")
			}
			currentElfCalories += currentInt
		}
	}

	fmt.Printf("Day 1 - Part 1: %d\n", mostCalories)
}

func partTwo() {
	file, err := os.Open("day01/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	allCalories := make([]int, 0)
	currentElfCalories := 0
	var currentLine string
	var currentInt int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine = scanner.Text()
		if currentLine == "" {
			allCalories = append(allCalories, currentElfCalories)
			currentElfCalories = 0
		} else {
			currentInt, err = strconv.Atoi(currentLine)
			if err != nil {
				log.Fatalf("Bad string to int conversion")
			}
			currentElfCalories += currentInt
		}
	}

	sort.Ints(allCalories)
	allCaloriesLen := len(allCalories)
	fmt.Printf("Day 1 - Part 2: %d\n", allCalories[allCaloriesLen-1]+allCalories[allCaloriesLen-2]+allCalories[allCaloriesLen-3])
}

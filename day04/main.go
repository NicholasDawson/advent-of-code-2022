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
	file, err := os.Open("day04/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	fullyContainCount := 0

	var line string
	var r1Start int
	var r1End int
	var r2Start int
	var r2End int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		splitLine := strings.Split(line, ",")
		r1 := strings.Split(splitLine[0], "-")
		r2 := strings.Split(splitLine[1], "-")

		r1Start, _ = strconv.Atoi(r1[0])
		r1End, _ = strconv.Atoi(r1[1])
		r2Start, _ = strconv.Atoi(r2[0])
		r2End, _ = strconv.Atoi(r2[1])

		if (r1Start <= r2Start && r1End >= r2End) || (r2Start <= r1Start && r2End >= r1End) {
			fullyContainCount++
		}
	}

	fmt.Printf("Day 4 - Part 1: %d\n", fullyContainCount)
}

func partTwo() {
	file, err := os.Open("day04/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	fullyContainCount := 0

	var line string
	var r1Start int
	var r1End int
	var r2Start int
	var r2End int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		splitLine := strings.Split(line, ",")
		r1 := strings.Split(splitLine[0], "-")
		r2 := strings.Split(splitLine[1], "-")

		r1Start, _ = strconv.Atoi(r1[0])
		r1End, _ = strconv.Atoi(r1[1])
		r2Start, _ = strconv.Atoi(r2[0])
		r2End, _ = strconv.Atoi(r2[1])

		if r2Start <= r1End && r1Start <= r2End {
			fullyContainCount++
		}
	}

	fmt.Printf("Day 4 - Part 2: %d\n", fullyContainCount)
}

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
	file, err := os.Open("day10/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	x := 1
	sumOfSignalStrengths := 0
	addCycle := 0

	var command string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for cycle := 1; cycle <= 220; cycle++ {
		command = scanner.Text()
		sumOfSignalStrengths += checkCycle(cycle) * x
		if addCycle == 1 {
			addx, _ := strconv.Atoi(strings.Split(command, " ")[1])
			x += addx
			addCycle = 0
			scanner.Scan()
		} else if command != "noop" {
			addCycle++
		} else {
			scanner.Scan()
		}
	}

	fmt.Printf("Day 9 - Part 1: %d\n", sumOfSignalStrengths)
}

func sliceContains(slice []int, num int) bool {
	for _, element := range slice {
		if element == num {
			return true
		}
	}
	return false
}

func checkCycle(cycle int) int {
	if sliceContains([]int{20, 60, 100, 140, 180, 220}, cycle) {
		return cycle
	}
	return 0
}

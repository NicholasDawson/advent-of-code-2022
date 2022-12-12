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

	fmt.Println("Day 10 - Part 2:")

	x := 1
	addCycle := 0

	var command string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Print("#") // First char always going to be lit
	for cycle := 1; cycle < 240; cycle++ {
		command = scanner.Text()
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

		// Draw pixel
		currentHorizontalPos := cycle % 40
		if currentHorizontalPos == 0 { // new line every 40
			fmt.Println()
		}

		if x-1 == currentHorizontalPos || x+1 == currentHorizontalPos || x == currentHorizontalPos {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
}

//     ###..###....##.#....####.#..#.#....###..
//     #..#.#..#....#.#....#....#..#.#....#..#.
//     ###..#..#....#.#....###..#..#.#....#..#.
//     #..#.###.....#.#....#....#..#.#....###..
//     #..#.#.#..#..#.#....#....#..#.#....#....
//     ###..#..#..##..####.#.....##..####.#....

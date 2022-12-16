package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day14/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(file)

	cave := &Cave{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		startX, startY := strToPoint(coords[0])
		for _, coord := range coords[1:] {
			x, y := strToPoint(coord)
			cave.addRockLine(startX, startY, x, y)
			startX, startY = x, y
		}
	}

	sandDropped := 0
	for ; cave.dropSand(); sandDropped++ {
		//cave.print(image.Point{X: 0, Y: 494}, image.Point{X: 9, Y: 503})
	}

	fmt.Printf("Day 14 - Part 1: %d\n", sandDropped)
}

func strToPoint(str string) (x, y int) {
	xAndY := strings.Split(str, ",")
	x, _ = strconv.Atoi(xAndY[0])
	y, _ = strconv.Atoi(xAndY[1])
	return x, y
}

const (
	AIR  uint8 = 0
	ROCK       = 1
	SAND       = 2
)

func displayFromEnum(enum uint8) string {
	switch enum {
	case AIR:
		return "."
	case ROCK:
		return "#"
	case SAND:
		return "O"
	default:
		return " "
	}
}

const dropX = 500
const gridMax = 1000

type Cave struct {
	grid [gridMax][gridMax]uint8
}

func (c *Cave) addRockLine(x1, y1, x2, y2 int) {
	if x1 == x2 {
		if y1 < y2 {
			for ; y1 <= y2; y1++ {
				c.grid[y1][x1] = ROCK
			}
		} else {
			for ; y2 <= y1; y2++ {
				c.grid[y2][x1] = ROCK
			}
		}
	} else {
		if x1 < x2 {
			for ; x1 <= x2; x1++ {
				c.grid[y1][x1] = ROCK
			}
		} else {
			for ; x2 <= x1; x2++ {
				c.grid[y1][x2] = ROCK
			}
		}
	}
}

func (c *Cave) dropSand() bool {
	x := dropX
	for y := 0; y < gridMax-1; y++ {
		below := c.grid[y+1][x]
		if below != AIR {
			if c.grid[y+1][x-1] == AIR {
				x--
			} else if c.grid[y+1][x+1] == AIR {
				x++
			} else {
				c.grid[y][x] = SAND
				return true
			}
		}
	}
	return false
}

func (c *Cave) print(min, max image.Point) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			fmt.Printf(displayFromEnum(c.grid[x][y]))
		}
		fmt.Println()
	}
}

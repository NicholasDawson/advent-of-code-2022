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

	cave.deepestY += 1
	cave.addRockLine(0, cave.deepestY, gridMax-1, cave.deepestY)
	sandDropped := 0
	for ; cave.dropSand(); sandDropped++ {
		//cave.print(image.Point{X: 0, Y: 490}, image.Point{X: 15, Y: 510})
	}

	fmt.Printf("Day 14 - Part 2: %d\n", sandDropped+1)
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
	grid     [gridMax][gridMax]uint8
	deepestY int
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

	if y1 > c.deepestY {
		c.deepestY = y1
	}
	if y2 > c.deepestY {
		c.deepestY = y2
	}
}

func (c *Cave) dropSand() bool {
	x := dropX
	for y := 0; y < c.deepestY; y++ {
		below := c.grid[y+1][x]
		if below != AIR {
			if c.grid[y+1][x-1] == AIR {
				x--
			} else if c.grid[y+1][x+1] == AIR {
				x++
			} else {
				c.grid[y][x] = SAND
				return c.grid[0][500] != SAND
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

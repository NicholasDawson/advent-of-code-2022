package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type RopeBridge struct {
	visitedTPositions [][2]int
	hx, hy            int
	tx, ty            int
}

func main() {
	file, err := os.Open("day9/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	ropeBridge := &RopeBridge{}

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		tokens := strings.Split(line, " ")
		direction := tokens[0]
		distance, _ := strconv.Atoi(tokens[1])

		switch direction {
		case "U":
			ropeBridge.up(distance)
		case "D":
			ropeBridge.down(distance)
		case "L":
			ropeBridge.left(distance)
		case "R":
			ropeBridge.right(distance)
		}

	}

	fmt.Printf("Day 9 - Part 1: %d\n", len(ropeBridge.visitedTPositions))
	fmt.Printf("Day 9 - Part 2: %d\n", 0)
}

func (rb *RopeBridge) visitedCoord(coord [2]int) bool {
	for _, element := range rb.visitedTPositions {
		if element == coord {
			return true
		}
	}
	return false
}
func (rb *RopeBridge) addCoordToVisited() {
	coord := [2]int{rb.tx, rb.ty}
	if !rb.visitedCoord(coord) {
		rb.visitedTPositions = append(rb.visitedTPositions, coord)
	}
}

func (rb *RopeBridge) adjustTail() {
	xDiff := rb.hx - rb.tx
	yDiff := rb.hy - rb.ty

	xDiffAbs := math.Abs(float64(xDiff))
	yDiffAbs := math.Abs(float64(yDiff))

	// Diagonal Adjustment
	if xDiffAbs == 2 && yDiffAbs == 1 {
		rb.tx += xDiff / 2
		rb.ty += yDiff
		println("Diagonal tail adjustment")
		rb.addCoordToVisited()
		return
	} else if xDiffAbs == 1 && yDiffAbs == 2 {
		rb.tx += xDiff
		rb.ty += yDiff / 2
		println("Diagonal tail adjustment")
		rb.addCoordToVisited()
		return
	}

	// Normal adjustment of tail keeping up
	println("Normal tail adjustment")
	if xDiff == 2 {
		rb.tx++
	} else if xDiff == -2 {
		rb.tx--
	} else if yDiff == 2 {
		rb.ty++
	} else if yDiff == -2 {
		rb.ty--
	}
	rb.addCoordToVisited()
}

func (rb *RopeBridge) print(msg string) {
	fmt.Printf("%s: H[%d, %d] T[%d, %d]\n", msg, rb.hx, rb.hy, rb.tx, rb.ty)
}

func (rb *RopeBridge) up(distance int) {
	for i := 0; i < distance; i++ {
		rb.hy++
		rb.adjustTail()
		rb.print("U")
	}
}
func (rb *RopeBridge) down(distance int) {
	for i := 0; i < distance; i++ {
		rb.hy--
		rb.adjustTail()
		rb.print("D")
	}
}
func (rb *RopeBridge) left(distance int) {
	for i := 0; i < distance; i++ {
		rb.hx--
		rb.adjustTail()
		rb.print("L")
	}
}
func (rb *RopeBridge) right(distance int) {
	for i := 0; i < distance; i++ {
		rb.hx++
		rb.adjustTail()
		rb.print("R")
	}
}

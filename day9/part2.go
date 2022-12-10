package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Rope struct {
	visitedTPositions []image.Point
	point             image.Point
	tail              *Rope
}

func main() {
	file, err := os.Open("day9/given2")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	// Create rope
	head := &Rope{}
	tail := head
	for i := 0; i < 9; i++ {
		newRope := &Rope{}
		tail.tail = newRope
		tail = newRope
	}

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		currentRope := head
		ropeIndex := 0
		for currentRope != nil {
			fmt.Printf("%d %s\n", ropeIndex, currentRope.point)
			currentRope = currentRope.tail
			ropeIndex++
		}

		fmt.Println()
		fmt.Println()
		fmt.Println(line)

		tokens := strings.Split(line, " ")
		direction := tokens[0]
		distance, _ := strconv.Atoi(tokens[1])

		switch direction {
		case "U":
			head.up(distance)
		case "D":
			head.down(distance)
		case "L":
			head.left(distance)
		case "R":
			head.right(distance)
		}
	}

	fmt.Printf("Day 9 - Part 2: %d\n", len(tail.visitedTPositions))
}

func (r *Rope) visitedPoint(point image.Point) bool {
	for _, element := range r.visitedTPositions {
		if element == point {
			return true
		}
	}
	return false
}
func (r *Rope) addPointToVisited() {
	newPoint := image.Point{
		X: r.point.X,
		Y: r.point.Y,
	}
	if !r.visitedPoint(newPoint) {
		r.visitedTPositions = append(r.visitedTPositions, newPoint)
	}
}

//func (r *Rope) adjustTail() {
//	if r.tail == nil { // Ignore last rope segment
//		r.addPointToVisited()
//		return
//	}
//
//	xDiff := r.point.X - r.tail.point.X
//	yDiff := r.point.Y - r.tail.point.Y
//
//	xDiffAbs := math.Abs(float64(xDiff))
//	yDiffAbs := math.Abs(float64(yDiff))
//
//	// Diagonal Adjustment
//	if xDiffAbs == 2 && yDiffAbs == 1 {
//		r.tail.point.X += xDiff / 2
//		r.tail.point.Y += yDiff
//		r.addPointToVisited()
//		r.tail.adjustTail()
//		return
//	} else if xDiffAbs == 1 && yDiffAbs == 2 {
//		r.tail.point.X += xDiff
//		r.tail.point.Y += yDiff / 2
//		r.addPointToVisited()
//		r.tail.adjustTail()
//		return
//	}
//
//	// Normal adjustment of tail keeping up
//	if xDiff == 2 {
//		r.tail.point.X++
//	} else if xDiff == -2 {
//		r.tail.point.X--
//	} else if yDiff == 2 {
//		r.tail.point.Y++
//	} else if yDiff == -2 {
//		r.tail.point.Y--
//	}
//	r.addPointToVisited()
//	r.tail.adjustTail()
//}

func (r *Rope) adjustTail() {
	if r.tail == nil { // Ignore last rope segment
		r.addPointToVisited()
		return
	}

	xDiff := r.point.X - r.tail.point.X
	yDiff := r.point.Y - r.tail.point.Y

	xDiffAbs := math.Abs(float64(xDiff))
	yDiffAbs := math.Abs(float64(yDiff))

	if xDiffAbs >= 2 || yDiffAbs >= 2 {
		r.tail.point.X = r.point.X
		r.tail.point.Y = r.point.Y
	}
	r.addPointToVisited()
	r.tail.adjustTail()
}

func (r *Rope) up(distance int) {
	for i := 0; i < distance; i++ {
		r.adjustTail()
		r.point.Y++
	}
}
func (r *Rope) down(distance int) {
	for i := 0; i < distance; i++ {
		r.adjustTail()
		r.point.Y--
	}
}
func (r *Rope) left(distance int) {
	for i := 0; i < distance; i++ {
		r.adjustTail()
		r.point.X--
	}
}
func (r *Rope) right(distance int) {
	for i := 0; i < distance; i++ {
		r.adjustTail()
		r.point.X++
	}
}

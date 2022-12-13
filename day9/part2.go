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

func main() {
	file, err := os.Open("day9/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file...", err)
		}
	}(file)

	// Create rope
	var rope []image.Point
	const ropeLength = 10
	for i := 0; i < ropeLength; i++ {
		rope = append(rope, image.Point{})
	}
	ropeLen := len(rope)

	// Keep track of positions the tail visited
	var visitedTailPositions []image.Point

	printRope(rope)

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)

		tokens := strings.Split(line, " ")
		direction := tokens[0]
		distance, _ := strconv.Atoi(tokens[1])

		switch direction {
		case "U":
			for i := 0; i < distance; i++ {
				adjustRope(ropeLen, rope, &visitedTailPositions, 0, 1)
				//printRope(rope)
			}
		case "D":
			for i := 0; i < distance; i++ {
				adjustRope(ropeLen, rope, &visitedTailPositions, 0, -1)
				//printRope(rope)
			}
		case "L":
			for i := 0; i < distance; i++ {
				adjustRope(ropeLen, rope, &visitedTailPositions, -1, 0)
				//printRope(rope)
			}
		case "R":
			for i := 0; i < distance; i++ {
				adjustRope(ropeLen, rope, &visitedTailPositions, 1, 0)
				//printRope(rope)
			}
		}
	}

	fmt.Printf("Day 9 - Part 2: %d\n", len(visitedTailPositions))
}

func visitPoint(point image.Point, visitedList *[]image.Point) {
	for _, element := range *visitedList {
		if element == point {
			return
		}
	}
	*visitedList = append(*visitedList, point)
}

func adjustRope(ropeLen int, rope []image.Point, visitedList *[]image.Point, xAdd, yAdd int) {
	// Move head
	rope[0].X += xAdd
	rope[0].Y += yAdd

	// Ensure we only adjust when the rope is at least 2 away
	headDiff := rope[0].Sub(rope[1])
	xDiffAbs := math.Abs(float64(headDiff.X))
	yDiffAbs := math.Abs(float64(headDiff.Y))
	if xDiffAbs <= 1 && yDiffAbs <= 1 {
		return
	}

	// Iterate through rest of points of rope and adjust them
	for index, point := range rope[1:] {
		index++ // Adjust index to account for removed head element

		pointDiff := rope[index-1].Sub(point)
		pointXDiffAbs := math.Abs(float64(pointDiff.X))
		pointYDiffAbs := math.Abs(float64(pointDiff.Y))

		// Check for diagonal movement
		if pointXDiffAbs <= 1 && pointYDiffAbs <= 1 {
			if index == ropeLen-1 {
				// Visit point for tail
				visitPoint(rope[index], visitedList)
			}
			return
		} else if pointXDiffAbs == 2 && pointYDiffAbs == 1 {
			rope[index].X += pointDiff.X / 2
			rope[index].Y += pointDiff.Y
		} else if pointXDiffAbs == 1 && pointYDiffAbs == 2 {
			rope[index].X += pointDiff.X
			rope[index].Y += pointDiff.Y / 2
		} else {
			if pointDiff.X == 2 {
				rope[index].X++
			} else if pointDiff.X == -2 {
				rope[index].X--
			}
			if pointDiff.Y == 2 {
				rope[index].Y++
			} else if pointDiff.Y == -2 {
				rope[index].Y--
			}
		}

		if index == ropeLen-1 {
			// Visit point for tail
			visitPoint(rope[index], visitedList)
		}
	}
}

func printRope(rope []image.Point) {
	// Display grid
	for y := 16 - 1; y >= -5; y-- {
		var covering []string
		for x := -11; x < 15; x++ {
			printedChar := false
			for knotIndex, knot := range rope {
				if knot.X == x && knot.Y == y {
					if !printedChar {
						if knotIndex == 0 {
							fmt.Print("H")
						} else {
							fmt.Print(knotIndex)
						}
						printedChar = true
					} else {
						covering = append(covering, strconv.Itoa(knotIndex))
					}
				}
			}
			if !printedChar {
				fmt.Print(".")
			}
		}
		if covering != nil {
			fmt.Printf("  (covering %s)", strings.Join(covering, ", "))
		}
		fmt.Println()
	}
	fmt.Println()

}

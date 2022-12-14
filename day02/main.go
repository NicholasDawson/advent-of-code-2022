package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

const (
	Lose = 0
	Tie  = 3
	Win  = 6
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	// FIRST COLUMN -> Opponent Play
	// A = ROCK
	// B = PAPER
	// C = SCISSORS

	// SECOND COLUMN -> Your Play
	// X = ROCK
	// Y = PAPER
	// Z = SCISSORS

	// Play points
	// 1 = ROCK
	// 2 = PAPER
	// 3 = SCISSORS

	// Outcome points
	// 0 = Lose
	// 3 = Tie
	// 6 = Win

	// Total score = sum of scores for each round
	// Score per round = played points + outcome points

	var playToType = map[uint8]int{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}

	file, err := os.Open("day02/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var currentLine string
	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		currentLine = scanner.Text()

		opponentPlay := playToType[currentLine[0]]
		yourPlay := playToType[currentLine[2]]

		totalScore += yourPlay // Score the shape you selected

		// Score outcome of match (do nothing for lose condition)
		if opponentPlay == yourPlay {
			totalScore += Tie
		} else if yourPlay == Rock && opponentPlay == Scissors {
			totalScore += Win
		} else if yourPlay == Paper && opponentPlay == Rock {
			totalScore += Win
		} else if yourPlay == Scissors && opponentPlay == Paper {
			totalScore += Win
		}
	}

	fmt.Printf("Day 2 - Part 1: %d\n", totalScore)
}

func partTwo() {
	var playOrOutcomeToType = map[uint8]int{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
		'X': Lose,
		'Y': Tie,
		'Z': Win,
	}

	var whatShapeBeatsMe = map[int]int{
		Rock:     Paper,
		Paper:    Scissors,
		Scissors: Rock,
	}

	var whatShapeLosesAgainstMe = map[int]int{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}

	file, err := os.Open("day02/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var currentLine string
	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		currentLine = scanner.Text()

		opponentPlay := playOrOutcomeToType[currentLine[0]]
		yourOutcome := playOrOutcomeToType[currentLine[2]]

		totalScore += yourOutcome // Score outcome

		// Score shape you had to pick for desired outcome
		if yourOutcome == Tie {
			totalScore += opponentPlay
		} else if yourOutcome == Win {
			totalScore += whatShapeBeatsMe[opponentPlay]
		} else if yourOutcome == Lose {
			totalScore += whatShapeLosesAgainstMe[opponentPlay]
		}
	}

	fmt.Printf("Day 2 - Part 2: %d\n", totalScore)
}

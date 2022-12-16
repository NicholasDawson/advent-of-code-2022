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

func main() {
	file, err := os.Open("day15/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(file)

	const chosenRow = 2000000
	positions := 0

	intervalList := &IntervalList{}
	var takenXPoints []int

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		tokens := strings.Split(line, "=")

		sensorX, _ := strconv.Atoi(tokens[1][:len(tokens[1])-3])
		sensorY, _ := strconv.Atoi(tokens[2][:len(tokens[2])-24])
		beaconX, _ := strconv.Atoi(tokens[3][:len(tokens[3])-3])
		beaconY, _ := strconv.Atoi(tokens[4])
		manhattenDistance := getManhattenDistance(sensorX, sensorY, beaconX, beaconY)

		if sensorY+manhattenDistance >= chosenRow && sensorY-manhattenDistance <= chosenRow {
			intervalList.addInterval(getInterval(sensorX, sensorY, chosenRow, manhattenDistance))
		}

		if sensorY == chosenRow {
			takenXPoints = append(takenXPoints, sensorX)
		}
		if beaconY == chosenRow && !sliceContains(takenXPoints, beaconX) {
			takenXPoints = append(takenXPoints, beaconX)
		}
	}

	positions -= len(takenXPoints)
	for _, interval := range intervalList.list {
		positions += interval.high - interval.low + 1
	}

	fmt.Printf("Day 15 - Part 1: %d\n", positions)
}

func getManhattenDistance(sensorX, sensorY, beaconX, beaconY int) int {
	return int(math.Abs(float64(sensorX-beaconX)) + math.Abs(float64(sensorY-beaconY)))
}

func getInterval(sensorX, sensorY, yLevel, manhattenDistance int) Interval {
	offset := manhattenDistance - int(math.Abs(float64(sensorY-yLevel)))
	return Interval{
		low:  sensorX - offset,
		high: sensorX + offset,
	}
}

func sliceContains(slice []int, num int) bool {
	for _, element := range slice {
		if element == num {
			return true
		}
	}
	return false
}

type Interval struct {
	low  int
	high int
}

type IntervalList struct {
	list []Interval
}

func (il *IntervalList) addInterval(newInterval Interval) {
	for index, interval := range il.list {
		// If overlaps replace existing interval with combined interval
		if newInterval.high >= interval.low && newInterval.low <= interval.high {
			if interval.high > newInterval.high {
				newInterval.high = interval.high
			}
			if interval.low < newInterval.low {
				newInterval.low = interval.low
			}
			il.list = append(il.list[:index], il.list[index+1:]...)
			il.addInterval(newInterval)
			return
		}
	}

	// Does not overlap add to list
	il.list = append(il.list, newInterval)
}

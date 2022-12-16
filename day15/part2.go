package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	entries := make([]Entry, 36)
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

		entries = append(entries, Entry{
			sensorX:           sensorX,
			sensorY:           sensorY,
			beaconX:           beaconX,
			beaconY:           beaconY,
			manhattenDistance: manhattenDistance,
		})
	}

	for y := 0; y <= 4000000; y++ {
		getRowTakenPositions(y, entries)
	}
}

func getRowTakenPositions(chosenRow int, entries []Entry) {
	//positions := 0
	intervalList := &IntervalList{}

	for _, entry := range entries {
		if entry.sensorY+entry.manhattenDistance >= chosenRow && entry.sensorY-entry.manhattenDistance <= chosenRow {
			intervalList.addInterval(getInterval(entry.sensorX, entry.sensorY, chosenRow, entry.manhattenDistance))
		}
	}

	//positions -= len(takenXPoints)
	//for _, interval := range intervalList.list {
	//	positions += interval.high - interval.low + 1
	//}

	if len(intervalList.list) >= 2 {
		sort.Slice(intervalList.list, func(i, j int) bool {
			return intervalList.list[i].low < intervalList.list[j].low
		})

		tuningFrequency := (intervalList.list[0].high+1)*4000000 + chosenRow
		fmt.Printf("Day 15 - Part 2: %v\n", tuningFrequency)
	}
	//return positions
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

type Entry struct {
	sensorX, sensorY  int
	beaconX, beaconY  int
	manhattenDistance int
}

type Interval struct {
	low  int
	high int
}

type IntervalList struct {
	list []Interval
}

func (il *IntervalList) addInterval(newInterval Interval) {
	if newInterval.low < 0 {
		newInterval.low = 0
	}
	if newInterval.high > 4000000 {
		newInterval.high = 4000000
	}

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

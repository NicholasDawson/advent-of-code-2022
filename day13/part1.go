package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	file, err := os.Open("day13/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(file)

	var line1, line2 string
	sumOfIndicesInRightOrder := 0
	scanner := bufio.NewScanner(file)
	for pairIndex := 1; scanner.Scan(); pairIndex++ {
		line1 = scanner.Text()
		scanner.Scan()
		line2 = scanner.Text()
		scanner.Scan() // Remove trailing new line

		packet1 := stringToList(line1)
		packet2 := stringToList(line2)

		fmt.Printf("== Pair %d ==\n", pairIndex)
		if isListValid(packet1, packet2) {
			sumOfIndicesInRightOrder += pairIndex
		}
	}

	fmt.Printf("Day 13 - Part 1: %d\n", sumOfIndicesInRightOrder)
}

func stringToList(text string) []any {
	list := make([]any, 0)
	if text == "[]" {
		return list
	}
	text = text[1 : len(text)-1]

	var currentToken string
	var openCount, closedCount int
	lastIndex := len(text) - 1
	for index, char := range text {
		if char == ',' {
			if openCount == 0 && closedCount == 0 {
				tokenAsInt, _ := strconv.Atoi(currentToken)
				list = append(list, tokenAsInt)
				currentToken = ""
				continue
			} else if openCount == closedCount {
				list = append(list, stringToList(currentToken))
				openCount, closedCount = 0, 0
				currentToken = ""
				continue
			}
		} else if char == '[' {
			openCount++
		} else if char == ']' {
			closedCount++
		}
		if index == lastIndex {
			if openCount == 0 && closedCount == 0 {
				currentToken += string(char)
				tokenAsInt, _ := strconv.Atoi(currentToken)
				list = append(list, tokenAsInt)
			} else if openCount == closedCount {
				currentToken += string(char)
				list = append(list, stringToList(currentToken))
			}
		}
		currentToken += string(char)
	}
	return list
}

func isListValid(list1, list2 []any) bool {
	fmt.Printf("- Compare %v vs %v\n", list1, list2)
	list2Len := len(list2)
	for index, element1 := range list1 {
		if index == list2Len {
			fmt.Println("\tRight side ran out of items, not right order")
			return false
		}
		element2 := list2[index]
		e1Kind := reflect.TypeOf(element1).Kind()
		e2Kind := reflect.TypeOf(element2).Kind()

		if e1Kind == reflect.Int && e2Kind == reflect.Int {
			fmt.Printf("\t- Compare %d vs %d\n", element1, element2)
			if element1.(int) > element2.(int) {
				fmt.Println("\t\t- Right side is smaller, so inputs are not in the right order")
				return false
			} else if element1.(int) < element2.(int) {
				fmt.Println("\t\t- Left side is smaller, so inputs are in the right order")
				return true
			}
		} else if e1Kind == reflect.Slice && e2Kind == reflect.Slice {
			if !isListValid(element1.([]any), element2.([]any)) {
				return false
			}
		} else if e1Kind == reflect.Int && e2Kind == reflect.Slice {
			if !isListValid([]any{element1.(int)}, element2.([]any)) {
				return false
			}
		} else if e1Kind == reflect.Slice && e2Kind == reflect.Int {
			if !isListValid(element1.([]any), []any{element2.(int)}) {
				return false
			}
		}
	}
	fmt.Println("\tLeft side ran out of items, right order")
	return true
}

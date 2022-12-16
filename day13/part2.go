package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day13/input2")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(file)

	var line string
	var packets [][]any
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			continue
		}

		packet := stringToList(line)
		packets = append(packets, packet)
	}

	sort.Slice(packets, func(i, j int) bool {
		return isListValid(packets[i], packets[j]) == 1
	})

	var divider1, divider2 int
	for index, packet := range packets {
		index++
		if reflect.DeepEqual(packet, []any{[]any{2}}) {
			divider1 = index
		} else if reflect.DeepEqual(packet, []any{[]any{6}}) {
			divider2 = index
		}
	}

	fmt.Printf("Day 13 - Part 2: %d\n", divider1*divider2)
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

func isListValid(list1, list2 []any) uint8 {
	fmt.Printf("- Compare %v vs %v\n", list1, list2)
	list1Len, list2Len := len(list1), len(list2)
	for index, element1 := range list1 {
		if index == list2Len {
			fmt.Println("\tRight side ran out of items, not right order")
			return 0
		}
		element2 := list2[index]
		e1Kind := reflect.TypeOf(element1).Kind()
		e2Kind := reflect.TypeOf(element2).Kind()

		if e1Kind == reflect.Int && e2Kind == reflect.Int {
			fmt.Printf("\t- Compare %d vs %d\n", element1, element2)
			if element1.(int) > element2.(int) {
				fmt.Println("\t\t- Right side is smaller, so inputs are not in the right order")
				return 0
			} else if element1.(int) < element2.(int) {
				fmt.Println("\t\t- Left side is smaller, so inputs are in the right order")
				return 1
			}
		} else if e1Kind == reflect.Slice && e2Kind == reflect.Slice {
			result := isListValid(element1.([]any), element2.([]any))
			if result != 2 {
				return result
			}
		} else if e1Kind == reflect.Int && e2Kind == reflect.Slice {
			result := isListValid([]any{element1.(int)}, element2.([]any))
			if result != 2 {
				return result
			}
		} else if e1Kind == reflect.Slice && e2Kind == reflect.Int {
			result := isListValid(element1.([]any), []any{element2.(int)})
			if result != 2 {
				return result
			}
		}
	}
	if list1Len == list2Len {
		return 2
	}
	fmt.Println("\tLeft side ran out of items, right order")
	return 1
}

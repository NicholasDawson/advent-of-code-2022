package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items           []int
	operator        string
	operand         int
	testDivisibleBy int
	testTrueMonkey  int
	testFalseMonkey int
	itemsInspected  int
}

func (m *Monkey) getMonkeyToThrowTo(newItem int) int {
	if newItem%m.testDivisibleBy == 0 {
		return m.testTrueMonkey
	} else {
		return m.testFalseMonkey
	}
}

func (m *Monkey) applyOperator(old int) int {
	return applyOperator(m.operator, m.operand, old)
}

func applyOperator(operator string, operand int, old int) int {
	// if operand is 0 then use old as operand
	if operand == 0 {
		operand = old
	}

	switch operator {
	case "+":
		return old + operand
	case "-":
		return old - operand
	case "*":
		return old * operand
	case "/":
		return old / operand
	default:
		return 0
	}
}

func main() {
	file, err := os.Open("day11/given")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var monkeys []Monkey

	// Read all monkeys into structs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Skip monkey title (Monkey 0:)
		newMonkey := Monkey{}

		// Get starting items
		scanner.Scan()
		startingItemsString := strings.Split(scanner.Text(), ": ")[1]
		startingItemNumbers := strings.Split(startingItemsString, ", ")
		for _, numAsString := range startingItemNumbers {
			numAsInt, _ := strconv.Atoi(numAsString)
			newMonkey.items = append(newMonkey.items, numAsInt)
		}

		// Get operator
		scanner.Scan()
		operationTokens := strings.Split(scanner.Text(), " ")
		newMonkey.operator = operationTokens[6]

		// Get operand
		operandString := operationTokens[7]
		if operandString != "old" { // old should be 0 so nothing necessary
			newMonkey.operand, _ = strconv.Atoi(operandString)
		}

		// Get Test divisible by
		scanner.Scan()
		testTokens := strings.Split(scanner.Text(), " ")
		newMonkey.testDivisibleBy, _ = strconv.Atoi(testTokens[5])

		// Get test true monkeys
		scanner.Scan()
		trueTestToken := strings.Split(scanner.Text(), " ")
		newMonkey.testTrueMonkey, _ = strconv.Atoi(trueTestToken[9])

		// Get test false monkeys
		scanner.Scan()
		falseTestToken := strings.Split(scanner.Text(), " ")
		newMonkey.testFalseMonkey, _ = strconv.Atoi(falseTestToken[9])

		// Read line for next monkey newline
		scanner.Scan()

		// Add monkey
		monkeys = append(monkeys, newMonkey)
	}

	monkeysLen := len(monkeys)

	for _, operator := range "+-*/" {
		fmt.Printf("Trying operator %s\n", string(operator))
		for i := -10000; i <= 10000; i++ {
			// fmt.Printf("%d\n", i)
			monkeysForIteration := make([]*Monkey, monkeysLen)
			for monkeyIndex, monkey := range monkeys {
				monkeyCopy := monkey
				monkeysForIteration[monkeyIndex] = &monkeyCopy
			}
			if monkeyPart2(monkeysForIteration, string(operator), i) {
				break
			}
		}
	}
	fmt.Println("nothing...")
}

func monkeyPart2(monkeys []*Monkey, operator string, operand int) bool {

	monkeyCount := len(monkeys)
	monkeyInspections := make([]int, monkeyCount)

	// Process 10000 turns
	for turn := 0; turn <= 1000; turn++ {
		for monkeyIndex, monkey := range monkeys {
			for _, item := range monkey.items {
				newItem := monkey.applyOperator(item)
				newItem = applyOperator(operator, operand, newItem)
				monkeyToThrowTo := monkey.getMonkeyToThrowTo(newItem)
				monkeys[monkeyToThrowTo].items = append(monkeys[monkeyToThrowTo].items, newItem)
				monkeyInspections[monkeyIndex]++
			}
			monkey.items = []int{} // clear current monkeys items after processing
		}
		//if (turn+1)%1000 == 0 || turn+1 == 1 || turn+1 == 20 {
		//	fmt.Printf("== After round %d ==\n", turn+1)
		//	for monkeyIndex, _ := range monkeys {
		//		fmt.Printf("Monkey %d inspected items %d times.\n", monkeyIndex, monkeyInspections[monkeyIndex])
		//	}
		//	fmt.Println()
		//}
		if turn+1 == 1000 {
			expectedValues := []int{5204, 4792, 199, 5192}
			for monkeyIndex := range monkeys {
				if expectedValues[monkeyIndex] != monkeyInspections[monkeyIndex] {
					return false
				}
			}
		}
	}

	sort.Ints(monkeyInspections)
	monkeyBusiness := monkeyInspections[monkeyCount-1] * monkeyInspections[monkeyCount-2]

	fmt.Printf("Day 11 - Part 2: %d\n", monkeyBusiness)
	return true
}

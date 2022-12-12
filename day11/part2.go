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

func (m *Monkey) getMonkeyToThrowTo(newItem int, debug bool) int {
	if newItem%m.testDivisibleBy == 0 {
		if debug {
			fmt.Printf("\t\tCurrent worry level is divisible by %d.\n", m.testDivisibleBy)
		}
		return m.testTrueMonkey
	} else {
		if debug {
			fmt.Printf("\t\tCurrent worry level is not divisible by %d.\n", m.testDivisibleBy)
		}
		return m.testFalseMonkey
	}
}

func (m *Monkey) applyOperator(old int, debug bool) int {
	// if operand is 0 then use old as operand
	operand := m.operand
	if operand == 0 {
		operand = old
	}

	switch m.operator {
	case "+":
		if debug {
			fmt.Printf("\t\tWorry level is increased by %d to %d.\n", operand, old+operand)
		}
		return old + operand
	case "-":
		if debug {
			fmt.Printf("\t\tWorry level is decreased by %d to %d.\n", operand, old-operand)
		}
		return old - operand
	case "*":
		if debug {
			fmt.Printf("\t\tWorry level is multiplied by %d to %d.\n", operand, old*operand)
		}
		return old * operand
	case "/":
		if debug {
			fmt.Printf("\t\tWorry level is divided by %d to %d.\n", operand, old/operand)
		}
		return old / operand
	default:
		log.Fatal("BAD")
		return 0
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcmArray(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcmArray(result, integers[i])
	}

	return result
}

func main() {
	const debug = false
	file, err := os.Open("day11/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var monkeys []*Monkey
	var monkeyDivideByInts []int

	// Read all monkeys into structs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Skip monkey title (Monkey 0:)
		newMonkey := &Monkey{}

		// Get starting items
		scanner.Scan()
		startingItemsString := strings.Split(scanner.Text(), ": ")[1]
		startingItemNumbers := strings.Split(startingItemsString, ", ")
		for _, numAsString := range startingItemNumbers {
			numAsInt, _ := strconv.Atoi(numAsString)
			newMonkey.items = append(newMonkey.items, numAsInt) // Prepend
		}

		// Get operator
		scanner.Scan()
		operationTokens := strings.Split(scanner.Text(), " ")
		newMonkey.operator = operationTokens[6]

		// Get operand
		operand := operationTokens[7]
		if operand != "old" { // old should be 0 so nothing necessary
			newMonkey.operand, _ = strconv.Atoi(operand)
		}

		// Get Test divisible by
		scanner.Scan()
		testTokens := strings.Split(scanner.Text(), " ")
		newMonkey.testDivisibleBy, _ = strconv.Atoi(testTokens[5])
		monkeyDivideByInts = append(monkeyDivideByInts, newMonkey.testDivisibleBy)

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

	lcmOfMonkeysDivision := lcmArray(monkeyDivideByInts[0], monkeyDivideByInts[1], monkeyDivideByInts[2:]...)
	monkeyCount := len(monkeys)
	monkeyInspections := make([]int, monkeyCount)

	// Process 1000 turns
	for turn := 0; turn < 10000; turn++ {
		for monkeyIndex, monkey := range monkeys {
			if debug {
				fmt.Printf("Monkey %d:\n", monkeyIndex)
			}
			for _, item := range monkey.items {
				if debug {
					fmt.Printf("\tMonkey inspects an item with a worry level of %d.\n", item)
				}
				newItem := monkey.applyOperator(item, debug)
				newItem %= lcmOfMonkeysDivision
				if debug {
					fmt.Printf("\t\tMonkey gets bored with item. Worry level is divided by 3 to %d.\n", newItem)
				}
				monkeyToThrowTo := monkey.getMonkeyToThrowTo(newItem, debug)
				if debug {
					fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d.\n", newItem, monkeyToThrowTo)
				}
				monkeys[monkeyToThrowTo].items = append(monkeys[monkeyToThrowTo].items, newItem) // Prepend
				monkeyInspections[monkeyIndex]++
			}
			monkey.items = []int{} // clear current monkeys items after processing
		}
		if debug {
			fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", turn+1)
			for monkeyIndex, monkey := range monkeys {
				fmt.Printf("Monkey %d: ", monkeyIndex)
				for _, item := range monkey.items {
					fmt.Printf("%d, ", item)
				}
				fmt.Printf("\n\tInspected itesm %d times.", monkeyInspections[monkeyIndex])
				fmt.Println()
			}
			fmt.Println()
		}
	}

	sort.Ints(monkeyInspections)
	monkeyBusiness := monkeyInspections[monkeyCount-1] * monkeyInspections[monkeyCount-2]

	fmt.Printf("Day 11 - Part 2: %d\n", monkeyBusiness)
}

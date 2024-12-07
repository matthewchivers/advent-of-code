package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	fmt.Println("Hello, advent of code 2024 - Day 7!")
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
}

func partOne(input []string) int {
	operators := []OperatorFunc{add, multiply}
	return calculateTotal(input, operators)
}

func partTwo(input []string) int {
	operators := []OperatorFunc{add, multiply, concatenate}
	return calculateTotal(input, operators)
}

// OperatorFunc defines the function signature for operators.
type OperatorFunc func(int, int) int

// Define operator functions
var (
	add         OperatorFunc = func(a, b int) int { return a + b }
	multiply    OperatorFunc = func(a, b int) int { return a * b }
	concatenate OperatorFunc = concatenateFunc
)

// calculateTotal processes the input lines and calculates the total result based on the testValue.
func calculateTotal(input []string, operators []OperatorFunc) int {
	var total int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Experimenting with go routines/"workers" to speed up the process
	numWorkers := 8
	jobs := make(chan string, len(input))

	// Define worker function
	// Each worker processes lines from the jobs channel and updates the total
	worker := func() {
		defer wg.Done()
		localTotal := 0
		for line := range jobs {
			if valid, testValue := processLine(line, operators); valid {
				localTotal += testValue
			}
		}
		mu.Lock()
		total += localTotal
		mu.Unlock()
	}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	// Send jobs to workers
	for _, line := range input {
		jobs <- line
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	return total
}

// processLine processes a single line and returns whether it's valid and its testValue
func processLine(line string, operators []OperatorFunc) (bool, int) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return false, 0
	}

	testValueStr := strings.TrimSpace(parts[0])
	testValue, err := strconv.Atoi(testValueStr)
	if err != nil {
		return false, 0
	}

	numbersStr := strings.TrimSpace(parts[1])
	if numbersStr == "" {
		return false, 0
	}
	numberParts := strings.Fields(numbersStr)
	numbers := make([]int, len(numberParts))
	for i, numStr := range numberParts {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return false, 0
		}
		numbers[i] = num
	}

	if len(numbers) == 1 {
		if numbers[0] == testValue {
			return true, testValue
		}
		return false, 0
	}

	// Calculate the total number of combinations of operators for the given input.
	// This is used to iterate through all possible combinations of operators.
	// For example, if there are 2 operators and 5 numbers, there are 2^4 = 16 combinations.

	opPositions := len(numbers) - 1
	totalCombinations := pow(len(operators), opPositions)
	valid := false

	for combo := 0; combo < totalCombinations && !valid; combo++ {
		result := numbers[0]
		currentCombo := combo

		// Loop through the positions between numbers and apply operators.
		// Use currentCombo to determine which operator to apply at each position by taking its remainder
		// when divided by the number of operators. This ensures all combinations are tried.
		// Divide currentCombo by the number of operators after each position to prepare for the next.

		for i := 0; i < opPositions; i++ {
			opIndex := currentCombo % len(operators)
			currentCombo /= len(operators)
			result = operators[opIndex](result, numbers[i+1])
			if result == -1 { // Early termination if concatenate fails
				break
			}
		}
		if result == testValue {
			valid = true
		}
	}

	if valid {
		return true, testValue
	}
	return false, 0
}

// concatenateFunc combines two integers by appending the digits of the second to the first.
// This avoids using string conversions for better performance.
func concatenateFunc(a, b int) int {
	if b == 0 {
		return a * 10
	}
	pow := 1
	tmp := b
	for tmp > 0 {
		pow *= 10
		tmp /= 10
	}
	return a*pow + b
}

// pow computes the power of an integer: base^exp.
// This is used to calculate the number of operator combinations for a given input length.
func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

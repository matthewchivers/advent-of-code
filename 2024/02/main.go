package main

import (
	"fmt"
	"math"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	lines = aoc.ReadFileAsLines("input.txt")
)

func main() {
	fmt.Println("Hello, advent of code 2024 - Day 2!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

func partOne() int {
	countSafe := 0
	for _, line := range lines {
		numbers, err := aoc.ParseIntList(line, " ")
		if err != nil {
			fmt.Println("Error converting string to int")
			return 0
		}
		if isReportSafe(numbers) {
			countSafe++
		}
	}
	return countSafe
}

func isReportSafe(report []int) bool {
	prevDiff := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		absVal := math.Abs(float64(diff))
		if absVal > 3 || absVal < 1 {
			return false
		}
		if diff*prevDiff < 0 {
			// if the sign of the difference changes (was positive but is now negative), return false
			return false
		}
		prevDiff = diff
	}
	return true
}

func partTwo() int {
	fmt.Println("Part two not implemented")
	return 0
}

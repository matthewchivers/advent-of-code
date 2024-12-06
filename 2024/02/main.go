package main

import (
	"fmt"
	"math"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

// Glossary:
// - "report" refers to a list of levels.
// - "level" refers to an individual integer value in the report.

func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	fmt.Println("Hello, advent of code 2024 - Day 2!")
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
}

// partOne counts how many reports are safe (adhere to rules)
func partOne(input []string) int {
	return countSafe(input, isReportSafe)
}

// partTwo counts how many reports are safe by removing up to one level.
func partTwo(input []string) int {
	return countSafe(input, isSafeWithOneRemoval)
}

// countSafe counts how many reports meet the provided safety criteria (safeFunc)
func countSafe(lines []string, safeFunc func([]int) bool) int {
	count := 0
	for _, line := range lines {
		numbers, err := aoc.ParseIntList(line, " ")
		if err == nil && safeFunc(numbers) {
			count++
		}
	}
	return count
}

// isReportSafe checks if a report is safe:
// - Differences between consecutive levels must be between 1 and 3.
// - The direction (increasing or decreasing) must be consistent.
func isReportSafe(report []int) bool {
	if len(report) <= 2 { // Reports with two or fewer levels are always safe
		return true
	}
	prevDiff := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		// Check if the difference is valid and direction does not alternate
		if abs := math.Abs(float64(diff)); abs > 3 || abs < 1 || (i > 0 && diff*prevDiff < 0) {
			return false
		}
		prevDiff = diff
	}
	return true
}

// isSafeWithOneRemoval checks if removing one level makes the report safe.
// Brute force (try all permutations) - presumably there is a more efficient way to do this?
func isSafeWithOneRemoval(report []int) bool {
	if isReportSafe(report) {
		return true
	}
	for i := range report {
		reducedReport := make([]int, len(report)-1)
		copy(reducedReport, report[:i])
		copy(reducedReport[i:], report[i+1:])
		if isReportSafe(reducedReport) {
			return true
		}
	}
	return false
}

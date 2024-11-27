package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	lines = aoc.ReadFileAsLines("input.txt")
)

func main() {
	// Solve and print the results for both parts of Day 4: Camp Cleanup
	log.Printf("Part One Total: %d", partOne())
	log.Printf("Part Two Total: %d", partTwo())
}

// partOne solves Part 1 of the Day 4 challenge by counting the number of times one range fully encompasses the other.
func partOne() int {
	return countMatchingRanges(isEncompassed)
}

// partTwo solves Part 2 of the Day 4 challenge by counting the number of times the ranges overlap.
func partTwo() int {
	return countMatchingRanges(isOverlapping)
}

// countMatchingRanges iterates over each line of input, extracting ranges and applying the given condition function.
// The conditionFunc determines which ranges are counted.
func countMatchingRanges(conditionFunc func([]int, []int) bool) int {
	count := 0
	for _, line := range lines {
		rangeOne, rangeTwo, err := getRanges(line)
		if err != nil {
			continue
		}
		if conditionFunc(rangeOne, rangeTwo) {
			count++
		}
	}
	return count
}

// getRanges takes a string input in the format "x-y,a-b" and returns two ranges represented as []int slices.
// It returns an error if the input format is incorrect or if the values cannot be parsed into integers.
func getRanges(s string) ([]int, []int, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid range format")
	}

	left := strings.Split(parts[0], "-")
	right := strings.Split(parts[1], "-")
	if len(left) != 2 || len(right) != 2 {
		return nil, nil, fmt.Errorf("invalid range format")
	}

	num1, err1 := strconv.Atoi(left[0])
	num2, err2 := strconv.Atoi(left[1])
	num3, err3 := strconv.Atoi(right[0])
	num4, err4 := strconv.Atoi(right[1])

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, nil, fmt.Errorf("invalid number conversion")
	}

	rangeOne := []int{num1, num2}
	rangeTwo := []int{num3, num4}

	return rangeOne, rangeTwo, nil
}

// isEncompassed checks if one range is fully contained within the other.
// It returns true if rangeOne is fully within rangeTwo or vice versa.
func isEncompassed(rangeOne, rangeTwo []int) bool {
	return (rangeOne[0] >= rangeTwo[0] && rangeOne[1] <= rangeTwo[1]) ||
		(rangeTwo[0] >= rangeOne[0] && rangeTwo[1] <= rangeOne[1])
}

// isOverlapping checks if there is any overlap between the two ranges.
// It returns true if any part of rangeOne overlaps with rangeTwo.
func isOverlapping(rangeOne, rangeTwo []int) bool {
	return !(rangeOne[1] < rangeTwo[0] || rangeOne[0] > rangeTwo[1])
}

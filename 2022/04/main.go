package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	lines = aoc.ReadFileAsLines("input.txt")
)

func main() {
	// Solve and print the results for both parts of the challenge
	log.Printf("Part One Total: %d", partOne())
	log.Printf("Part Two Total: %d", partTwo())
}

// partOne is a wrapper function that solves part one of the challenge by counting the number of times one range fully encompasses the other.
func partOne() int {
	return countMatchingRanges(isEncompassed)
}

// partTwo is a wrapper function that solves part two of the challenge by counting the number of times the ranges overlap.
func partTwo() int {
	return countMatchingRanges(isOverlapping)
}

// countMatchingRanges iterates over each line of input, extracting ranges and applying the given condition function.
// The conditionFunc determines which ranges are counted.
func countMatchingRanges(conditionFunc func([2]int, [2]int) bool) int {
	count := 0
	for _, line := range lines {
		// Extract the ranges from the line
		rangeOne, rangeTwo, err := getRanges(line)
		if err != nil {
			// Log and skip lines that do not match the expected format
			log.Printf("Skipping invalid line: %s", line)
			continue
		}

		// Apply the given condition function
		if conditionFunc(rangeOne, rangeTwo) {
			count++
		}
	}
	return count
}

// getRanges takes a string input in the format "x-y,a-b", and returns two ranges represented as [2]int arrays.
// It returns an error if the input format is incorrect or if the values cannot be parsed into integers.
func getRanges(s string) ([2]int, [2]int, error) {
	// Compile a regular expression to match the pattern we want: two ranges in the format "x-y,a-b"
	re := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	matches := re.FindStringSubmatch(s)

	// Ensure we have exactly five matches (full match + four capture groups)
	if len(matches) != 5 {
		// Return an error if the input doesn't match the expected format
		return [2]int{}, [2]int{}, fmt.Errorf("invalid range format")
	}

	// Convert the matched strings to integers
	num1, err1 := strconv.Atoi(matches[1])
	num2, err2 := strconv.Atoi(matches[2])
	num3, err3 := strconv.Atoi(matches[3])
	num4, err4 := strconv.Atoi(matches[4])

	// Check for errors during conversion
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return [2]int{}, [2]int{}, fmt.Errorf("invalid number conversion")
	}

	// Create the ranges from the parsed numbers
	rangeOne := [2]int{num1, num2}
	rangeTwo := [2]int{num3, num4}

	return rangeOne, rangeTwo, nil
}

// isEncompassed checks if one range is fully encompassed by the other.
// It returns true if rangeOne is fully within rangeTwo or vice versa.
func isEncompassed(rangeOne, rangeTwo [2]int) bool {
	// Return true if rangeOne is fully within rangeTwo or vice versa
	return (rangeOne[0] >= rangeTwo[0] && rangeOne[1] <= rangeTwo[1]) ||
		(rangeTwo[0] >= rangeOne[0] && rangeTwo[1] <= rangeOne[1])
}

// isOverlapping checks if there is any overlap between the two ranges.
// It returns true if any part of rangeOne overlaps with rangeTwo.
func isOverlapping(rangeOne, rangeTwo [2]int) bool {
	// Return true if the ranges overlap at any point
	return !(rangeOne[1] < rangeTwo[0] || rangeOne[0] > rangeTwo[1])
}

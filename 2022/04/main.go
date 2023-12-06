package main

import (
	"log"
	"regexp"
	"strconv"

	aoc "github.com/matthewchivers/advent-of-code/utils"
)

var (
	lines = aoc.ReadFileAsLines("input.txt")
)

func main() {
	log.Println("Part One Total:", solvePart(1))
	log.Println("Part Two Total:", solvePart(2))
}

func solvePart(part int) int {
	count := 0
	for _, line := range lines {
		rangeOne, rangeTwo := getRanges(line)
		if (part == 1 && checkEncompassment(rangeOne, rangeTwo)) ||
			(part == 2 && checkOverlap(rangeOne, rangeTwo)) {
			count++
		}
	}
	return count
}

func getRanges(s string) ([2]int, [2]int) {
	// Compile a regular expression to match the pattern we want
	re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

	matches := re.FindStringSubmatch(s)
	num1, _ := strconv.Atoi(matches[1])
	num2, _ := strconv.Atoi(matches[2])
	num3, _ := strconv.Atoi(matches[3])
	num4, _ := strconv.Atoi(matches[4])

	slice1 := [2]int{num1, num2}
	slice2 := [2]int{num3, num4}

	return slice1, slice2
}

func checkEncompassment(rangeOne, rangeTwo [2]int) bool {
	// Checking if [87 87] encompasses [6 86]
	if (rangeOne[0] >= rangeTwo[0] && rangeOne[1] <= rangeTwo[1]) ||
		(rangeTwo[0] >= rangeOne[0] && rangeTwo[1] <= rangeOne[1]) {
		return true
	}
	return false
}

func checkOverlap(rangeOne, rangeTwo [2]int) bool {
	// Checking if [87 87] overlaps [6 86]
	if rangeOne[1] < rangeTwo[0] {
		return false
	}
	if rangeOne[0] > rangeTwo[1] {
		return false
	}
	return true
}

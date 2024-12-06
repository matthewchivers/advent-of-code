package main

import (
	"fmt"
	"strings"

	"slices"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	fmt.Println("Hello, advent of code 2024 - Day 1!")
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
}

// partOne calculates the total distance between paired location IDs from two lists
func partOne(input []string) int {
	listA, listB := parseLists(input)
	slices.Sort(listA)
	slices.Sort(listB)

	totalDiff := 0
	for i := range listA {
		totalDiff += abs(listA[i] - listB[i])
	}
	return totalDiff
}

// partTwo calculates the similarity score based on frequency of matching location IDs
func partTwo(input []string) int {
	listA, listB := parseLists(input)

	// Create a frequency map for listB
	freqMap := make(map[int]int)
	for _, num := range listB {
		freqMap[num]++
	}

	similarityScore := 0
	for _, num := range listA {
		count := freqMap[num]
		similarityScore += num * count
	}

	return similarityScore
}

// parseLists extracts two lists of integers from input lines
func parseLists(input []string) ([]int, []int) {
	listA, listB := []int{}, []int{}
	for _, line := range input {
		parts := strings.Split(line, "   ")
		intA, _ := aoc.StringToInt(parts[0])
		intB, _ := aoc.StringToInt(parts[1])
		listA = append(listA, intA)
		listB = append(listB, intB)
	}
	return listA, listB
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

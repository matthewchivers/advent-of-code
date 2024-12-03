package main

import (
	"fmt"
	"strings"

	"slices"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var lines = aoc.ReadFileAsLines("input.txt")

func main() {
	fmt.Println("Hello, advent of code 2024 - Day 1!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// partOne calculates the total distance between paired location IDs from two lists
func partOne() int {
	listA, listB := parseLists()
	slices.Sort(listA)
	slices.Sort(listB)

	totalDiff := 0
	for i := range listA {
		totalDiff += abs(listA[i] - listB[i])
	}
	return totalDiff
}

// partTwo calculates a similarity score based on occurrences of shared location IDs
func partTwo() int {
	occurrences := map[int]int{}
	listA, listB := parseLists()

	for _, intA := range listA {
		occurrences[intA] = 0
	}

	total := 0
	for _, num := range listB {
		if _, ok := occurrences[num]; ok {
			occurrences[num]++
		}
	}

	for k, v := range occurrences {
		total += k * v
	}
	return total
}

// parseLists extracts two lists of integers from input lines
func parseLists() ([]int, []int) {
	listA, listB := []int{}, []int{}
	for _, line := range lines {
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

package main

import (
	"fmt"
	"strings"

	"slices"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	lines = aoc.ReadFileAsLines("input.txt")
)

func main() {
	fmt.Println("Hello, advent of code 2024 - Day 1!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// partOne calculates the total distance between paired location IDs from two lists
func partOne() int {
	listA := []int{}
	listB := []int{}

	// Parse the input lines to populate listA and listB
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		intA, err := aoc.StringToInt(parts[0])
		if err != nil {
			fmt.Println("Error converting string to int")
			return 0
		}
		listA = append(listA, intA)
		intB, err := aoc.StringToInt(parts[1])
		if err != nil {
			fmt.Println("Error converting string to int")
			return 0
		}
		listB = append(listB, intB)
	}

	// Sort both lists to ensure proper pairing of elements
	slices.Sort(listA)
	slices.Sort(listB)

	countDiff := 0
	// Calculate the absolute difference for each pair and add to the total
	for i := 0; i < len(listA); i++ {
		diff := listA[i] - listB[i]
		if diff < 0 {
			diff = -diff
		}
		countDiff += diff
	}
	return countDiff
}

// partTwo calculates a similarity score based on occurrences of shared location IDs
func partTwo() int {
	listB := []int{}
	occurences := map[int]int{}

	// Parse the input lines to populate occurrences and listB
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		intA, err := aoc.StringToInt(parts[0])
		if err != nil {
			fmt.Println("Error converting string to int")
			return 0
		}
		occurences[intA] = 0

		intB, err := aoc.StringToInt(parts[1])
		if err != nil {
			fmt.Println("Error converting string to int")
			return 0
		}
		listB = append(listB, intB)
	}

	// Count the occurrences of numbers in listB that also appear in the occurrences map
	for _, num := range listB {
		_, ok := occurences[num]
		if ok {
			occurences[num]++
		}
	}

	rangeTotal := 0
	// Calculate the similarity score by multiplying each unique value by its frequency
	for k, v := range occurences {
		rangeTotal += k * v
	}

	return rangeTotal
}

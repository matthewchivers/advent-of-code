package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

var (
	topElves      []int
	lines         []string
	inputFileName = "../input.txt"
)

func main() {
	log.Println("Part 1: Highest number of calories being carried by a single elf: ", calculateCaloriesTopElves(1))
	log.Println("Part 2: Total number of calories being carried by the top three elves: ", calculateCaloriesTopElves(3))
}

// Calculates the total calories held by the top n elves
func calculateCaloriesTopElves(n int) int {
	populatePodium(n)
	if n == 1 {
		return topElves[0]
	}
	total := 0
	for i := 0; i < n; i++ {
		total += topElves[i]
	}
	return total
}

// Populates the topElves slice based on data from the input file
func populatePodium(n int) {
	topElves = make([]int, n)
	currentCaloriesHeld := 0
	for _, line := range aoc.ReadLines(inputFileName) {
		if line != "" {
			currentCaloriesHeld += aoc.StringToInt(line)
		} else {
			placeElf(currentCaloriesHeld)
			currentCaloriesHeld = 0
		}
	}
}

// Helper function to place an elf in the correct position in the topElves slice
func placeElf(val int) {
	lowestValueIndex := 0
	if val > topElves[lowestValueIndex] {
		topElves[lowestValueIndex] = val
		sortElf(lowestValueIndex)
	}
}

// Sorts the most recently added elf into the correct position in the topElves slice
// Recursively calls itself until the elf is in the correct position
// i = position of the elf to be sorted
func sortElf(i int) {
	prevIndex := i - 1
	if prevIndex >= 0 && topElves[prevIndex] > topElves[i] {
		topElves[prevIndex], topElves[i] = topElves[i], topElves[prevIndex]
	}
	if i < (len(topElves) - 1) {
		sortElf(i + 1)
	}
	return
}

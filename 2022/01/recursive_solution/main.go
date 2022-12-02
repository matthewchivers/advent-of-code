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
	log.Println("Part 1: Highest number of calories being carried by a single elf: ", calculatePartOne())
	log.Println("Part 2: Total number of calories being carried by the top three elfs: ", calculatePartTwo())
}

func calculatePartOne() int {
	return calculateTopN(1)
}

func calculatePartTwo() int {
	return calculateTopN(3)
}

func calculateTopN(n int) int {
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

func populatePodium(n int) {
	topElves = make([]int, n)
	lines = aoc.ReadLines(inputFileName)
	currentCaloriesHeld := 0
	for _, line := range lines {
		if line != "" {
			currentCaloriesHeld += aoc.StringToInt(line)
		} else {
			placeElf(currentCaloriesHeld)
			currentCaloriesHeld = 0
		}
	}
}

func placeElf(val int) {
	if val > topElves[0] {
		topElves[0] = val
		sortElf(0)
	}
}

func sortElf(i int) {
	prevIndex := i - 1
	if prevIndex >= 0 {
		if topElves[prevIndex] > topElves[i] {
			topElves[prevIndex], topElves[i] = topElves[i], topElves[prevIndex]
		}
	}
	if i < (len(topElves) - 1) {
		sortElf(i + 1)
	}

	return
}

package main

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/advent-of-code/utils"
)

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 4!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// partOne returns the answer to part one of the day's puzzle
// Parse score cards (lottery tickets of sorts) each winning number on a card doubles the score (score of 1 for the first match)
// Returns the total score for all cards
func partOne() int {
	lines := utils.ReadFileAsString("input.txt")
	score := 0
	for _, line := range lines {
		winners, entries := parseLine(line)
		lineScore := 0
		for _, entry := range entries {
			if _, ok := winners[entry]; ok {
				if lineScore == 0 {
					lineScore = 1
				} else {
					lineScore *= 2
				}
			}
		}
		score += lineScore
	}
	return score
}

func parseLine(line string) (map[string]bool, []string) {
	// Card   1: 10  5 11 65 27 43 44 29 24 69 | 65 66 18 14 17 97 95 34 38 23 10 25 22 15 87  9 28 43  4 71 89 20 72  5  6
	colonIndex := strings.Index(line, ":")
	line = line[colonIndex+2:]
	split := strings.Split(line, "|")
	winners := strings.Fields(split[0])
	winnersMap := make(map[string]bool)
	for _, winner := range winners {
		winnersMap[winner] = true
	}
	entries := strings.Fields(split[1])
	return winnersMap, entries
}

// partTwo returns the answer to part two of the day's puzzle
func partTwo() int {
	return 0
}

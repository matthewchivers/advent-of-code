package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

// main is the entry point of the program.
func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	log.Println("Part One Total:", partOne(lines))
	log.Println("Part Two Total:", partTwo(lines))
}

// partOne calculates the total value for part one by splitting each line and finding the common character.
func partOne(input []string) int {
	total := 0
	for _, line := range input {
		firstHalf, secondHalf := splitStringInHalf(line)
		commonRune := findCommonRune(firstHalf, secondHalf)
		total += getRuneValue(commonRune)
	}
	return total
}

// partTwo calculates the total value for part two by finding the common character across chunks of three lines.
func partTwo(input []string) int {
	totalSum := 0
	for i := 0; i < len(input); i += 3 {
		commonRune := findCommonRune(input[i], input[i+1], input[i+2])
		totalSum += getRuneValue(commonRune)
	}
	return totalSum
}

// getRuneValue returns the value of a given rune. Lowercase letters are 1-26, uppercase letters are 27-52.
func getRuneValue(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}

// findCommonRune finds the first common rune across all input strings.
func findCommonRune(toCompare ...string) rune {
	runeCount := make(map[rune]int)
	for _, str := range toCompare {
		seen := make(map[rune]bool)
		for _, r := range str {
			if !seen[r] {
				runeCount[r]++
				seen[r] = true
			}
			if runeCount[r] == len(toCompare) {
				return r
			}
		}
	}
	return 0
}

// splitStringInHalf splits the given string into two equal halves.
func splitStringInHalf(line string) (string, string) {
	firstHalf, secondHalf := line[:len(line)/2], line[len(line)/2:]
	return firstHalf, secondHalf
}

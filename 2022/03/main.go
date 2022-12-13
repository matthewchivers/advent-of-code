package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

var (
	lines = aoc.ReadLinesAsString("input.txt")
)

func main() {
	log.Println("Part One Total:", partOne())
	log.Println("Part Two Total:", partTwo())
}

func partOne() int {
	total := 0
	for _, line := range lines {
		firstHalf, secondHalf := splitStringInHalf(line)
		duplicate := getMatch(firstHalf, secondHalf)
		total += getDupeValue(duplicate)
	}
	return total
}

func partTwo() int {
	totalSum := 0
	// iterate over three lines at a time
	for i := 0; i < len(lines); i += 3 {
		match := getMatch(lines[i], lines[i+1], lines[i+2])
		totalSum += getDupeValue(match)
	}
	return totalSum
}

func getDupeValue(dupe rune) int {
	if dupe > 96 {
		return (int(dupe) - 96)
	}
	return (int(dupe) - 38)
}

func getMatch(toCompare ...string) rune {
	runeCount := make(map[rune]int)
	for _, str := range toCompare {
		alreadyCounted := []rune{}
		for _, r := range str {
			if !aoc.SliceContainsRune(alreadyCounted, r) {
				runeCount[rune(r)]++
				alreadyCounted = append(alreadyCounted, r)
			}
			if runeCount[rune(r)] == len(toCompare) {
				return r
			}
		}
	}
	return 0
}

func splitStringInHalf(line string) (string, string) {
	firstHalf, secondHalf := line[:len(line)/2], line[len(line)/2:]
	return firstHalf, secondHalf
}

package main

import (
	"log"
	"strconv"
	"strings"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

var (
	lines = aoc.ReadLines("input.txt")
)

func main() {
	log.Println("Part One Total:", partOne())
}

func partOne() int {
	count := 0
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		if checkEncompassment(getRange(ranges[0]), getRange(ranges[1])) {
			count++
		}
	}
	return count
}

func getRange(rangeStr string) [2]int {
	ranges := strings.Split(rangeStr, "-")
	valOne, err := strconv.Atoi(ranges[0])
	if err != nil {
		log.Fatal(err)
	}
	valTwo, err := strconv.Atoi(ranges[1])
	if err != nil {
		log.Fatal(err)
	}
	return [2]int{valOne, valTwo}
}

func checkEncompassment(range1, range2 [2]int) bool {
	log.Println("Checking if", range1, "encompasses", range2, "(or vice versa)")
	// Checking if [87 87] encompasses [6 86]
	if (range1[0] >= range2[0] && range1[1] <= range2[1]) ||
		(range2[0] >= range1[0] && range2[1] <= range1[1]) {
		log.Println("Encompassment found")
		return true
	}
	log.Println("Encompassment not found")
	return false
}

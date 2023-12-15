package main

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/advent-of-code/utils"
)

var (
	numberMap = map[string]string{
		"one": "1", "two": "2", "three": "3",
		"four": "4", "five": "5", "six": "6",
		"seven": "7", "eight": "8", "nine": "9",
	}
)

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 1!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

func partOne() int {
	return loopLinesAndChars(1)
}

func partTwo() int {
	return loopLinesAndChars(2)
}

func loopLinesAndChars(part int) int {
	sum := 0
	lines := utils.ReadFileAsLines("input.txt")
	for _, line := range lines {
		var first, last string
		for i, character := range line {
			detectedChar := ""
			if character >= '0' && character <= '9' {
				detectedChar = string(character)
			}
			if part == 2 {
				for word, number := range numberMap {
					if strings.HasPrefix(line[i:], word) {
						detectedChar = number
						break
					}
				}
			}
			if detectedChar != "" {
				if first == "" {
					first = string(detectedChar)
				}
				last = string(detectedChar)
			}
		}
		if integer, err := utils.StringToInt(first + last); err != nil {
			fmt.Println(err)
		} else {
			sum += integer
		}
	}
	return sum
}

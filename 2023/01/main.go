package main

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/advent-of-code/util"
)

var (
	numberMap = map[string]string{
		"one": "1", "two": "2", "three": "3",
		"four": "4", "five": "5", "six": "6",
		"seven": "7", "eight": "8", "nine": "9",
	}
	startsWith = map[byte][]string{}
	endsWith   = map[byte][]string{}
)

func init() {
	for word := range numberMap {
		startsWith[word[0]] = append(startsWith[word[0]], word)
		endsWith[word[len(word)-1]] = append(endsWith[word[len(word)-1]], word)
	}
}

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 1!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// partOne returns the sum of the first and last numbers (digits) in each line
func partOne() int {
	return loopLinesAndChars(1)
}

// partTwo returns the sum of the first and last numbers (digits or words) in each line
func partTwo() int {
	return loopLinesAndChars(2)
}

// loopLinesAndChars loops through the lines and characters of the input file
func loopLinesAndChars(part int) int {
	sum := 0
	lines := util.ReadFileAsLines("input.txt")
	for _, line := range lines {
		first, last := extractFirstLast(line, part)
		if integer, err := util.StringToInt(first + last); err != nil {
			fmt.Println(err)
		} else {
			sum += integer
		}
	}
	return sum
}

// extractFirstLast extracts the first and last number from a string
// part 1: numbers are digits ("1", "2", "3")
// part 2: numbers are digits ("1", "2", "3") or words ("one", "two", "three")
func extractFirstLast(line string, part int) (string, string) {
	first, last := "", ""
	for iLeft, iRight := 0, len(line)-1; iLeft <= len(line) && iRight >= 0; iLeft, iRight = iLeft+1, iRight-1 {
		if first == "" {
			if line[iLeft] >= '0' && line[iLeft] <= '9' {
				first = string(line[iLeft])
			} else if part == 2 {
				if words, ok := startsWith[line[iLeft]]; ok {
					for _, word := range words {
						if strings.HasPrefix(line[iLeft:], word) {
							first = numberMap[word]
						}
					}
				}
			}
		}
		if last == "" {
			if line[iRight] >= '0' && line[iRight] <= '9' {
				last = string(line[iRight])
			} else if part == 2 {
				if words, ok := endsWith[line[iRight]]; ok {
					for _, word := range words {
						if strings.HasSuffix(line[:iRight+1], word) {
							last = numberMap[word]
						}
					}
				}
			}
		}
		if first != "" && last != "" {
			break
		}
	}
	return first, last
}

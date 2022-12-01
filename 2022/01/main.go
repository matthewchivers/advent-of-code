package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

var (
	podium Podium
)

func main() {
	lines := aoc.ReadLines("input.txt")

	populatePodium(lines)

	log.Println("Part 1: Highest number of calories being carried by a single elf: ", getPartOneResult())
	log.Println("Part 2: Total number of calories being carried by the top three elfs: ", getPartTwoResult())
}

// Returns the correct answer for part one
func getPartOneResult() int {
	return podium.Highest()
}

// Returns the correct answer for part two
func getPartTwoResult() int {
	return podium.Total()
}

// Takes lines of input and passes the totals (for each elf) to a Podium
func populatePodium(lines []string) {
	currentCaloriesHeld := 0
	for _, line := range lines {
		if line != "" {
			currentCaloriesHeld += aoc.StringToInt(line)
		} else {
			// Reached end of current elf
			podium.Insert(currentCaloriesHeld)
			currentCaloriesHeld = 0
		}
	}
}

// Podium stores the top three values passed to it
type Podium struct {
	first, second, third int
}

// Insert a new value into the "podium"
func (p *Podium) Insert(val int) {
	if val > p.third {
		p.third = val
		if p.third > p.second {
			p.third, p.second = p.second, p.third
			if p.second > p.first {
				p.second, p.first = p.first, p.second
			}
		}
	}
}

// Highest returns the highest value stored in Podium
func (p *Podium) Highest() int {
	return p.first
}

// Total returns the sum of the top three values stored in Podium
func (p *Podium) Total() int {
	return p.first + p.second + p.third
}

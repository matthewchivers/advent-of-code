package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	podium    Podium
	populated bool
	lines     []string
)

func main() {
	log.Println("Part 1: Highest number of calories being carried by a single elf: ", calculatePartOne())
	log.Println("Part 2: Total number of calories being carried by the top three elfs: ", calculatePartTwo())
}

// Returns the correct answer for part one
func calculatePartOne() int {
	populatePodium(lines)
	return podium.Highest()
}

// Returns the correct answer for part two
func calculatePartTwo() int {
	populatePodium(lines)
	return podium.Total()
}

// Takes lines of input and passes the totals (for each elf) to a Podium
func populatePodium(lines []string) {
	if populated {
		return
	}
	if len(lines) == 0 {
		lines = aoc.ReadFileAsLines("input.txt")
	}
	currentCaloriesHeld := 0
	for _, line := range lines {
		if line != "" {
			cals, err := aoc.StringToInt(line)
			if err != nil {
				log.Fatal(err)
			}
			currentCaloriesHeld += cals
		} else {
			// Reached end of current elf
			podium.Insert(currentCaloriesHeld)
			currentCaloriesHeld = 0
		}
	}
	populated = true
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

package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

func main() {
	lines := aoc.ReadLines("../input.txt")

	var podium = Podium{}
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

	log.Println("Top Three Total: ", podium.Total())
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

// Total returns the sum of the top three values stored in Podium
func (p *Podium) Total() int {
	return p.first + p.second + p.third
}

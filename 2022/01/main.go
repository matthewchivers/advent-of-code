package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	// Print results for both parts of the challenge
	log.Println("Part 1: Highest number of calories being carried by a single elf: ", calculatePartOne(lines))
	log.Println("Part 2: Total number of calories being carried by the top three elves: ", calculatePartTwo(lines))
}

// calculatePartOne calculates and returns the highest number of calories being carried by a single elf.
func calculatePartOne(input []string) int {
	podium := populatePodium(input)
	return podium.Highest()
}

// calculatePartTwo calculates and returns the total number of calories carried by the top three elves.
func calculatePartTwo(input []string) int {
	podium := populatePodium(input)
	return podium.Total()
}

// populatePodium populates the Podium struct with calorie totals for each elf.
// This function reads the input lines and processes each line, summing up calorie counts for each elf
// and adding them to the Podium.
func populatePodium(lines []string) Podium {
	podium := Podium{}
	currentCaloriesHeld := 0

	for _, line := range lines {
		if line != "" {
			// Convert the line to an integer representing calories
			cals, err := aoc.StringToInt(line)
			if err != nil {
				log.Fatal(err) // Terminate if an error occurs while converting string to integer
			}
			currentCaloriesHeld += cals
		} else {
			// Empty line indicates end of current elf's calorie count
			podium.Insert(currentCaloriesHeld)
			currentCaloriesHeld = 0
		}
	}
	// Handle the last elf if there's no trailing empty line
	podium.Insert(currentCaloriesHeld)

	return podium
}

// Podium stores the top three highest calorie counts among all elves
// It maintains first, second, and third highest values
type Podium struct {
	first, second, third int
}

// Insert takes a new calorie count and updates the top three values in the Podium.
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

// Highest returns the highest value stored in the Podium (i.e., the most calories carried by any single elf).
func (p *Podium) Highest() int {
	return p.first
}

// Total returns the sum of the top three values stored in the Podium.
// This represents the total number of calories being carried by the top three elves.
func (p *Podium) Total() int {
	return p.first + p.second + p.third
}

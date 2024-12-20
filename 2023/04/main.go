package main

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/advent-of-code/util"
)

type ScratchCard struct {
	winningNumbers map[string]bool
	entryNumbers   []string
	count          int
}

func NewScratchCard(winningNumbers map[string]bool, entryNumbers []string) *ScratchCard {
	return &ScratchCard{
		winningNumbers: winningNumbers,
		entryNumbers:   entryNumbers,
		count:          1,
	}
}

func main() {
	lines := util.ReadFileAsLines("input.txt")
	fmt.Println("Hello, advent of code 2023 - Day 4!")
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
}

// partOne returns the answer to part one of the day's puzzle
// Each winning number on a card doubles the score (score 1 for the first match)
// Returns the total score for all cards
func partOne(input []string) int {
	score := 0
	for _, scratchCard := range getInventory(input) {
		lineScore := 0
		for _, entry := range scratchCard.entryNumbers {
			if scratchCard.winningNumbers[entry] {
				lineScore += util.MaxInt(lineScore, 1)
			}
		}
		score += lineScore
	}
	return score
}

// partTwo returns the answer to part two of the day's puzzle
// Each winning number on a card grants a free card to the next card in the inventory
// e.g. 4 matches grants 1 copy each of the next 4 cards
// This stacks (if I have 2 cards with 4 matches, I get 2 copies each of the next 4 cards)
// Returns the total number of cards
func partTwo(input []string) int {
	inv := getInventory(input)
	for cardNumber := 1; cardNumber < len(inv)+1; cardNumber++ {
		matches := 0
		scratchcard := *inv[cardNumber]
		for _, entry := range scratchcard.entryNumbers {
			if scratchcard.winningNumbers[entry] {
				matches++
				nextIndex := cardNumber + matches
				if _, ok := inv[nextIndex]; ok {
					inv[nextIndex].count += scratchcard.count
				}
			}
		}
	}

	totalCards := 0
	for _, card := range inv {
		totalCards += card.count
	}
	return totalCards
}

func parseLine(line string) (map[string]bool, []string) {

	lineAfterColon := strings.SplitN(line, ": ", 2)[1]
	parts := strings.Split(lineAfterColon, " | ")

	winningNumbers := strings.Fields(parts[0])
	entryNumbers := strings.Fields(parts[1])

	// Convert winning numbers to a map for fast lookup
	winningNumbersMap := make(map[string]bool)
	for _, winner := range winningNumbers {
		winningNumbersMap[winner] = true
	}

	return winningNumbersMap, entryNumbers
}

func getInventory(input []string) map[int]*ScratchCard {
	inv := make(map[int]*ScratchCard)
	for i, line := range input {
		// i + 1 because the card numbers start at 1
		// makes debugging more human-friendly
		inv[i+1] = NewScratchCard(parseLine(line))
	}
	return inv
}

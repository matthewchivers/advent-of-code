package main

import (
	"bytes"
	"log"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

func main() {
	points := []byte("ABCXYZ")
	lines := aoc.ReadLines("input.txt")
	totalScore := 0
	for _, line := range lines {
		opponentSelection := bytes.IndexByte(points, line[0])%3 + 1
		mySelection := bytes.IndexByte(points, line[2])%3 + 1
		totalScore += getScore(mySelection, opponentSelection)
	}
	log.Println("Total score:", totalScore)

}

func getScore(mySelection, opponentSelection int) int {
	score := 0
	if mySelection == opponentSelection {
		score += 3 // draw
	} else if mySelection == 1 && opponentSelection == 3 ||
		mySelection == 2 && opponentSelection == 1 ||
		mySelection == 3 && opponentSelection == 2 {
		score += 6 // win
	} // else lose
	score += mySelection
	return score
}

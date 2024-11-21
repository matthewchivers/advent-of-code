package main

import (
	"bytes"
	"log"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	points                = []byte("ABCXYZ")                 // Maps enemy instructions (A, B, C) and player instructions (X, Y, Z)
	lines                 = aoc.ReadFileAsLines("input.txt") // Reads the input file as lines
	rock, paper, scissors = 1, 2, 3                          // Assigns numeric values to rock, paper, and scissors
)

func main() {
	log.Println("Total Score - Part One:", processLines(1))
	log.Println("Total Score - Part Two:", processLines(2))
}

// processLines processes each line of input based on the problem's part (1 or 2).
// It calculates and returns the total score.
func processLines(part int) int {
	score := 0
	for _, line := range lines {
		// enemyInstruction and myInstruction are derived from the input line.
		enemyInstruction := bytes.IndexByte(points, line[0])%3 + 1
		myInstruction := bytes.IndexByte(points, line[2])%3 + 1
		// Different scoring mechanisms are used depending on the part of the problem.
		if part == 1 {
			getScorePartOne(&enemyInstruction, &myInstruction, &score)
		} else {
			getScorePartTwo(&enemyInstruction, &myInstruction, &score)
		}
	}
	return score
}

// getScorePartOne calculates the score for part one of the problem.
// - Draws add 3 points.
// - Wins add 6 points plus the value of the move.
// - Losses add only the value of the move.
func getScorePartOne(enemyInstruction, myInstruction, score *int) {
	if *myInstruction == *enemyInstruction {
		*score += 3 // Draw condition
	} else if *myInstruction == 1 && *enemyInstruction == 3 ||
		*myInstruction == 2 && *enemyInstruction == 1 ||
		*myInstruction == 3 && *enemyInstruction == 2 {
		*score += 6 // Win condition
	} // No additional points for losing
	*score += *myInstruction // Add the value of the move
}

// getScorePartTwo calculates the score for part two of the problem.
// - Depending on the player's instruction, the outcome could be a win, draw, or loss.
// - Scores are adjusted accordingly, adding fixed points for the outcome and the move.
func getScorePartTwo(enemyInstruction, myInstruction, score *int) {
	switch *myInstruction {
	case 1:
		// I need to lose, so pick the move that loses to the enemy's choice.
		switch *enemyInstruction {
		case rock:
			*score += scissors
		case paper:
			*score += rock
		case scissors:
			*score += paper
		}
	case 2:
		// I need to draw, add draw points and match the enemy's move.
		*score += 3
		*score += *enemyInstruction
	case 3:
		// I need to win, add win points and pick the winning move.
		*score += 6
		switch *enemyInstruction {
		case rock:
			*score += paper
		case paper:
			*score += scissors
		case scissors:
			*score += rock
		}
	}
}

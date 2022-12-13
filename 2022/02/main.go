package main

import (
	"bytes"
	"log"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

var (
	points                = []byte("ABCXYZ")
	lines                 = aoc.ReadLinesAsString("input.txt")
	rock, paper, scissors = 1, 2, 3
)

func main() {
	log.Println("Total Score - Part One:", processLines(1))
	log.Println("Total Score - Part Two:", processLines(2))
}

func processLines(part int) int {
	score := 0
	for _, line := range lines {
		log.Println("Line:", string(line))
		enemyInstruction := bytes.IndexByte(points, line[0])%3 + 1
		myInstruction := bytes.IndexByte(points, line[2])%3 + 1
		if part == 1 {
			getScorePartOne(&enemyInstruction, &myInstruction, &score)
		} else {
			getScorePartTwo(&enemyInstruction, &myInstruction, &score)
		}
	}
	return score
}

func getScorePartOne(enemyInstruction, myInstruction, score *int) {
	if *myInstruction == *enemyInstruction {
		*score += 3 // draw
	} else if *myInstruction == 1 && *enemyInstruction == 3 ||
		*myInstruction == 2 && *enemyInstruction == 1 ||
		*myInstruction == 3 && *enemyInstruction == 2 {
		*score += 6 // win
	} // else lose
	*score += *myInstruction
}

func getScorePartTwo(enemyInstruction, myInstruction, score *int) {
	switch *myInstruction {
	case 1:
		// I need to lose
		switch *enemyInstruction {
		case rock:
			*score += scissors
		case paper:
			*score += rock
		case scissors:
			*score += paper
		}
	case 2:
		// I need to draw
		*score += 3
		*score += *enemyInstruction
	case 3:
		// I need to win
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

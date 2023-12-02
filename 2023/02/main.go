package main

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/advent-of-code/utils"
)

// GameResult is a struct to hold the results of a game
type GameResult struct {
	red   int
	green int
	blue  int
}

var (
	gameMax = GameResult{red: 12, green: 13, blue: 14}
)

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 2!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// partOne returns the answer to part one of the day's puzzle
// Add all GameIDs that have a red, green and blue value less than or equal to the gameMax
func partOne() int {
	sum := 0
	lines := utils.ReadFileAsString("input.txt")
	for i, line := range lines {
		gameRes := processGameLine(line)
		if gameRes.red <= gameMax.red && gameRes.green <= gameMax.green && gameRes.blue <= gameMax.blue {
			sum += i + 1 // GameID starts at 1 and continues sequentially
		}
	}
	return sum
}

// partTwo returns the answer to part two of the day's puzzle
// Add all the power sets of red, green and blue for each game
func partTwo() int {
	sum := 0
	lines := utils.ReadFileAsString("input.txt")
	for _, line := range lines {
		gameRes := processGameLine(line)
		sum += gameRes.red * gameRes.green * gameRes.blue
	}
	return sum
}

// Gets the maximum red, green and blue values for each game
// Example Line: "Game 3: 4 green, 1 blue; 6 blue, 5 green, 1 red; 11 green, 10 blue"
func processGameLine(line string) GameResult {

	// remove "Game \d: " from the start of the line
	colInd := strings.Index(line, ":")
	game := line[colInd+2:]
	draws := strings.Split(game, ";")

	gameRes := GameResult{red: 0, green: 0, blue: 0}
	for _, draw := range draws {
		ballSet := strings.Split(strings.TrimSpace(draw), ", ")
		for _, colourGroup := range ballSet {
			colourCount := strings.Split(colourGroup, " ")
			number, err := utils.StringToInt(colourCount[0])
			if err != nil {
				fmt.Println(err)
			}
			colour := colourCount[1]
			switch colour {
			case "red":
				gameRes.red = utils.MaxInt(gameRes.red, number)
			case "green":
				gameRes.green = utils.MaxInt(gameRes.green, number)
			case "blue":
				gameRes.blue = utils.MaxInt(gameRes.blue, number)
			}
		}
	}
	return gameRes
}

package main

import (
	"fmt"

	"github.com/matthewchivers/advent-of-code/utils"
)

var directions = []struct{ dx, dy int }{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1}, // up, down, left, right
	{1, 1}, {-1, -1}, {1, -1}, {-1, 1}, // diagonals
}

type coord struct {
	x int
	y int
}

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 3!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// partOne returns the answer to part one of the day's puzzle
// The puzzle is to find the sum of all numbers that have a symbol adjacent to them
func partOne() int {
	matrix := utils.ReadFileAsRuneMatrix("input.txt")

	sum := 0

	for y := 0; y < len(matrix); y++ {
		numberBuffer := ""
		includeNumber := false
		for x := 0; x < len(matrix[y]); x++ {
			character := matrix[x][y]
			if character >= '0' && character <= '9' {
				numberBuffer += string(character)
				includeNumber = includeNumber || hasAdjacentSymbol(matrix, x, y)
			}
			if (character > '9' || character < '0' || x == len(matrix)-1) && numberBuffer != "" {
				numInt, _ := utils.StringToInt(numberBuffer)
				if includeNumber {
					sum += numInt
				}
				numberBuffer = ""
				includeNumber = false
			}
		}
	}

	return sum
}

// partTwo returns the answer to part two of the day's puzzle
// The puzzle is to add the powers of two numbers that are adjacent to the same gear
func partTwo() int {
	matrix := utils.ReadFileAsRuneMatrix("input.txt")

	sum := 0

	gearMap := map[coord][]int{}
	for y := 0; y < len(matrix); y++ {
		numberBuffer := ""
		gears := make(map[coord]bool)
		for x := 0; x < len(matrix[y]); x++ {
			character := matrix[x][y]
			if character >= '0' && character <= '9' {
				numberBuffer += string(character)
				for _, gear := range getAdjacentGears(matrix, x, y) {
					gears[gear] = true
				}
			}
			if (character > '9' || character < '0' || x == len(matrix)-1) && numberBuffer != "" {
				numInt, _ := utils.StringToInt(numberBuffer)
				if len(gears) > 0 {
					for gear := range gears {
						gearMap[gear] = append(gearMap[gear], numInt)
					}
				}
				numberBuffer = ""
				gears = make(map[coord]bool)
			}
		}
	}

	for _, gear := range gearMap {
		if len(gear) == 2 {
			sum += gear[0] * gear[1]
		}
	}

	return sum
}

func hasAdjacentSymbol(matrix [][]rune, x int, y int) bool {
	//check if the given x and y has a symbol adjacent to it
	for _, direction := range directions {
		newX, newY := x+direction.x, y+direction.y
		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[x]) {
			character := matrix[newX][newY]
			if isSymbol(character) {
				return true
			}
		}
	}
	return false
}

func getAdjacentGears(matrix [][]rune, x int, y int) []coord {
	// identifies all adjacent gears to the given x and y
	gears := []coord{}
	for _, direction := range directions {
		newX, newY := x+direction.dx, y+direction.dy
		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[x]) {
			character := matrix[newX][newY]
			charString := string(character)
			if charString == "*" {
				gears = append(gears, coord{newX, newY})
			}
		}
	}
	return gears
}

func isSymbol(character rune) bool {
	// return true if character is not a number or a '.'
	return (character < '0' || character > '9') && character != '.'
}

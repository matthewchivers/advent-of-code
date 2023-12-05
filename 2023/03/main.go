package main

import (
	"fmt"

	"github.com/matthewchivers/advent-of-code/utils"
)

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 3!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// partOne returns the answer to part one of the day's puzzle
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
				includeNumber = includeNumber || checkSurroundingCharacters(matrix, x, y)
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
func partTwo() int {
	matrix := utils.ReadFileAsRuneMatrix("input.txt")

	sum := 0

	for y := 0; y < len(matrix); y++ {
		numberBuffer := ""
		includeNumber := false
		for x := 0; x < len(matrix[y]); x++ {
			character := matrix[x][y]
			if character >= '0' && character <= '9' {
				numberBuffer += string(character)
				// isSymbol, isGear := checkSurroundingCharacters(matrix, x, y)
				includeNumber = includeNumber || checkSurroundingCharacters(matrix, x, y)
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

func checkSurroundingCharacters(matrix [][]rune, x int, y int) bool {
	//check if the character is a number or a '.'
	// surrounding characters are any character up, down, left, right, or diagonal from the given x and y
	// if the character is not a number or a '.', then it is a surrounding character

	directions := []struct{ dx, dy int }{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1}, // up, down, left, right
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1}, // diagonals
	}
	isSymbol := false
	// isGear := false
	for _, direction := range directions {
		newX, newY := x+direction.dx, y+direction.dy
		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[x]) {
			character := matrix[newX][newY]
			charString := string(character)
			if isValidCharacter(character) && charString != "" {
				isSymbol = true
				// if charString == "*" {
				// 	isGear = true
				// }
			}
		}
	}

	return isSymbol
}

func isValidCharacter(character rune) bool {
	// return true if character is not a number or a '.'
	return (character < '0' || character > '9') && character != '.'
}

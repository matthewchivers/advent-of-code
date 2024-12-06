package main

import (
	"fmt"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	directions = []struct {
		dx, dy int
	}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
)

func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	fmt.Println("Hello, advent of code 2024 - Day 4!")
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
}

// partOne searches for occurrences of the word "XMAS" in the grid in all possible directions.
// It counts (and returns) the number of times the word appears.
func partOne(input []string) int {
	word := "XMAS"
	count := 0
	grid := input
	// Loop over each position in the grid.
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			// If the character matches the first letter of the word, search in all directions.
			if grid[y][x] == word[0] {
				for _, dir := range directions {
					if matchWord(grid, x, y, dir.dx, dir.dy, word) {
						count++
					}
				}
			}
		}
	}
	return count
}

// partTwo checks for the specific pattern "A" surrounded by "MAS" or "SAM" diagonally in the grid.
// It's an "X-MAS" (MAS/SAM in an X shape) pattern. "of course".
// partTwo counts (and returns) the number of times such a pattern is found.
func partTwo(input []string) int {
	grid := input
	count := 0
	// Loop over each position in the grid.
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			// If the character is 'A', check if it forms a valid "XMAS" pattern.
			if grid[y][x] == 'A' {
				if isXMas(grid, x, y) {
					count++
				}
			}
		}
	}
	return count
}

// matchWord checks if a given word can be found in the grid starting at (startX, startY) in a specified direction (dx, dy).
func matchWord(grid []string, startX, startY, dx, dy int, word string) bool {
	x, y := startX, startY
	for i := 0; i < len(word); i++ {
		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
			return false
		}
		if grid[y][x] != word[i] {
			return false
		}
		x += dx
		y += dy
	}
	return true
}

// isXMas checks if there is an "A" at (cx, cy) surrounded by "MAS" or "SAM" diagonally in the grid.
func isXMas(grid []string, cx, cy int) bool {
	// Ensure the central character is 'A'.
	if !inGrid(grid, cx, cy) || grid[cy][cx] != 'A' {
		return false
	}

	// Check top-left to bottom-right diagonal for "MAS" or "SAM".
	d1 := getDiagonalChars(grid, cx, cy, -1, -1, 1, 1) // chars for diagonal1
	// Check top-right to bottom-left diagonal for "MAS" or "SAM".
	d2 := getDiagonalChars(grid, cx, cy, 1, -1, -1, 1) // chars for diagonal2

	// If either diagonal is empty (out of bounds), return false.
	if d1 == "" || d2 == "" {
		return false
	}

	// Both diagonals must be "MAS" or "SAM" to be valid.
	return isMasLine(d1) && isMasLine(d2)
}

// getDiagonalChars returns a string consisting of the three characters forming a diagonal around (cx, cy).
// The function calculates the coordinates for the first and last characters in the diagonal.
func getDiagonalChars(grid []string, cx, cy, dx1, dy1, dx2, dy2 int) string {
	// First character of the diagonal.
	x1, y1 := cx+dx1, cy+dy1
	// Last character of the diagonal.
	x2, y2 := cx+dx2, cy+dy2
	// If either character is out of bounds, return an empty string.
	if !inGrid(grid, x1, y1) || !inGrid(grid, x2, y2) {
		return ""
	}
	// Return the concatenation of the first character, the centre character, and the last character.
	return string(grid[y1][x1]) + string(grid[cy][cx]) + string(grid[y2][x2])
}

// inGrid checks if the coordinates (x, y) are within the bounds of the grid.
func inGrid(grid []string, x, y int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y])
}

// isMasLine checks if a given string is either "MAS" or "SAM".
func isMasLine(s string) bool {
	return s == "MAS" || s == "SAM"
}

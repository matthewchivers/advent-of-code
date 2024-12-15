package main

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/advent-of-code/util"
)

func main() {
	lines := util.ReadFileAsLines("input.txt")
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines)) // No Part Two for this problem, placeholder.
}

// partOne calculates the number of unique positions visited by the guard.
func partOne(input []string) int {
	grid, guardPos, guardDir := parseInput(input)
	directions := map[rune][2]int{'^': {-1, 0}, 'v': {1, 0}, '<': {0, -1}, '>': {0, 1}}

	visited := make(map[[2]int]bool)

	height, width := len(grid), len(grid[0])
	for {
		// Mark the current position as visited
		visited[guardPos] = true

		// Calculate the next position based on the current direction
		nextRow, nextCol := guardPos[0]+directions[guardDir][0], guardPos[1]+directions[guardDir][1]

		// Stop if the next position is outside the grid
		if nextRow < 0 || nextRow >= height || nextCol < 0 || nextCol >= width {
			break
		}

		// If the next cell is a wall, turn right; otherwise, move forward
		if grid[nextRow][nextCol] == '#' {
			guardDir = turnRight(guardDir)
		} else {
			guardPos = [2]int{nextRow, nextCol}
		}
	}

	// Return the number of unique positions visited
	return len(visited)
}

// partTwo calculates the number of cells that would cause the guard to enter a loop if blocked.
func partTwo(input []string) int {
	grid, guardPos, guardDir := parseInput(input)
	height, width := len(grid), len(grid[0])
	startPos := guardPos

	count := 0
	// loop through each cell in the grid, to check if adding an obstacle there causes a loop
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			// Skip the guard's starting position and non-free cells
			if (r == startPos[0] && c == startPos[1]) || grid[r][c] != '.' {
				continue
			}

			// Check if placing an obstacle at (r, c) causes a loop
			if causesLoop(grid, guardPos, guardDir, [2]int{r, c}) {
				count++
			}
		}
	}

	// Return the count of cells causing a loop
	return count
}

// causesLoop checks if placing an obstacle causes the guard to enter a loop.
func causesLoop(grid []string, guardPos [2]int, guardDir rune, obstacle [2]int) bool {
	directions := map[rune][2]int{'^': {-1, 0}, 'v': {1, 0}, '<': {0, -1}, '>': {0, 1}}

	visitedStates := make(map[[3]int]bool)
	height, width := len(grid), len(grid[0])
	dirToIndex := map[rune]int{'^': 0, '>': 1, 'v': 2, '<': 3}

	pos, dir := guardPos, guardDir
	for {
		// Create a unique key for the current state (position and direction)
		stateKey := [3]int{pos[0], pos[1], dirToIndex[dir]}
		// If the state has been seen before, a loop is detected
		if visitedStates[stateKey] {
			return true
		}
		visitedStates[stateKey] = true

		// Calculate the next position based on the current direction
		nextPos := [2]int{pos[0] + directions[dir][0], pos[1] + directions[dir][1]}

		// Stop if the next position is outside the grid
		if nextPos[0] < 0 || nextPos[0] >= height || nextPos[1] < 0 || nextPos[1] >= width {
			return false
		}

		// If the next position is the obstacle or a wall, turn right
		if nextPos == obstacle || grid[nextPos[0]][nextPos[1]] == '#' {
			dir = turnRight(dir)
			continue
		}

		// Move to the next position
		pos = nextPos
	}
}

// parseInput parses the grid input and extracts the guard's position and direction.
func parseInput(input []string) ([]string, [2]int, rune) {
	var guardPos [2]int
	var guardDir rune
	grid := make([]string, len(input))

	for i, line := range input {
		grid[i] = line
		// Find the guard's initial position and direction in the grid
		if idx := strings.IndexAny(line, "^v<>"); idx != -1 {
			guardPos = [2]int{i, idx}
			guardDir = rune(line[idx])
		}
	}

	return grid, guardPos, guardDir
}

// turnRight rotates the guard's direction 90 degrees clockwise.
func turnRight(direction rune) rune {
	return map[rune]rune{'^': '>', '>': 'v', 'v': '<', '<': '^'}[direction]
}

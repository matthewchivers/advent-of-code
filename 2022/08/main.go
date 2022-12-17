package main

import (
	"fmt"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

func main() {
	fmt.Println("Part 1:", partOne())
}

func partOne() int {
	grid := aoc.ReadIntoGrid("input.txt")
	return processGrid(grid)
}

// Checks every tree on the grid to see if it is visible
func processGrid(grid [][]int) int {
	visibleTrees := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			isVisible := processCoordinate(grid, i, j)
			if isVisible {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func processCoordinate(grid [][]int, x, y int) bool {

	vertical := checkXY(grid, x, y, "y")
	horizontal := checkXY(grid, x, y, "x")
	return (horizontal || vertical)
}

func checkXY(grid [][]int, x, y int, direction string) bool {
	horizontal := checkXandY("x", grid, x, y)
	vertical := checkXandY("y", grid, x, y)
	return (horizontal || vertical)
}

func checkXandY(direction string, grid [][]int, x, y int) bool {
	fixedPointN := 0
	ceiling := 0
	if direction == "x" {
		fixedPointN = x
		ceiling = len(grid[0])
	} else {
		fixedPointN = y
		ceiling = len(grid)
	}
	height := grid[x][y]
	preN, postN := true, true
	// check pre-N
	for i := (fixedPointN - 1); i >= 0; i-- {
		surroundingTreeHeight := 0
		if direction == "x" {
			surroundingTreeHeight = grid[i][y]
		} else {
			surroundingTreeHeight = grid[x][i]
		}
		if surroundingTreeHeight >= height {
			preN = false
			break
		}
	}
	// check post-N
	for i := (fixedPointN + 1); i < ceiling; i++ {
		surroundingTreeHeight := 0
		if direction == "x" {
			surroundingTreeHeight = grid[i][y]
		} else {
			surroundingTreeHeight = grid[x][i]
		}
		if surroundingTreeHeight >= height {
			postN = false
			break
		}
	}
	return (preN || postN)
}

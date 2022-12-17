package main

import (
	"fmt"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

type coord struct {
	height      int
	x, y        int
	scenicValue int
	isVisible   bool
}

type grid [][]coord

func (g *grid) readIntoGrid(fileName string) {
	basicGrid := aoc.ReadIntoGrid("input.txt")
	*g = make([][]coord, len(basicGrid))
	for i := 0; i < len(basicGrid); i++ {
		(*g)[i] = make([]coord, len(basicGrid[i]))
		for j := 0; j < len(basicGrid[i]); j++ {
			(*g)[i][j].height = basicGrid[i][j]
			(*g)[i][j].x = i
			(*g)[i][j].y = j
		}
	}
}

func (g *grid) processGrid() {
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			g.processCoord(i, j)
		}
	}
}

func (g *grid) processCoord(x, y int) {
	(*g)[x][y].scenicValue = 1
	horizontal := g.checkXorY("x", x, y)
	vertical := g.checkXorY("y", x, y)
	(*g)[x][y].isVisible = (horizontal || vertical)
}

func (g *grid) checkXorY(direction string, x, y int) bool {
	fixedPointN := 0
	ceiling := 0
	if direction == "x" {
		fixedPointN = x
		ceiling = len((*g)[0])
	} else {
		fixedPointN = y
		ceiling = len((*g))
	}
	preN, postN := true, true
	// check pre-N
	scenicPre := 0
	for i := (fixedPointN - 1); i >= 0; i-- {
		surroundingTreeHeight := 0
		if direction == "x" {
			surroundingTreeHeight = (*g)[i][y].height
		} else {
			surroundingTreeHeight = (*g)[x][i].height
		}
		if surroundingTreeHeight >= (*g)[x][y].height {
			scenicPre++
			preN = false
			break
		}
		scenicPre++
	}
	// check post-N
	scenicPost := 0
	for i := (fixedPointN + 1); i < ceiling; i++ {
		surroundingTreeHeight := 0
		if direction == "x" {
			surroundingTreeHeight = (*g)[i][y].height
		} else {
			surroundingTreeHeight = (*g)[x][i].height
		}
		if surroundingTreeHeight >= (*g)[x][y].height {
			scenicPost++
			postN = false
			break
		}
		scenicPost++
	}
	(*g)[x][y].scenicValue *= scenicPre * scenicPost
	return (preN || postN)
}

func (g *grid) countVisibleTrees() int {
	visibleTrees := 0
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if (*g)[i][j].isVisible {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func (g *grid) getHighestScenicValue() coord {
	scenicTree := coord{}
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if (*g)[i][j].scenicValue > scenicTree.scenicValue {
				scenicTree = (*g)[i][j]
			}
		}
	}
	return scenicTree
}

func main() {
	fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo().scenicValue)
}

func partOne() int {
	forestMap := grid{}
	forestMap.readIntoGrid("input.txt")
	forestMap.processGrid()
	return forestMap.countVisibleTrees()
}

func partTwo() coord {
	forestMap := grid{}
	forestMap.readIntoGrid("input.txt")
	forestMap.processGrid()
	return forestMap.getHighestScenicValue()
}

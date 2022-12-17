package adventofcode

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadLinesAsString reads a file line by line and returns a slice of strings
func ReadLinesAsString(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	return lines
}

// ReadFileAsBytes reads a file and returns a slice of bytes
func ReadFileAsBytes(fileName string) []byte {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return dat
}

// 30373   << size 5
// 25512   \/ size 6
// 65332
// 33549
// 35390
// 12345

// ReadIntoGrid reads a file line by line and returns a 2D slice of ints
func ReadIntoGrid(fileName string) [][]int {

	// x ->
	// 30373  y
	// 25512 ||
	// 65332 \/

	// grid[x][y] is the ideal setup
	// grid[] needs to be the length of x
	// grid[x] needs to be the length of y
	lines := ReadLinesAsString(fileName)
	grid := [][]int{}
	for y, line := range lines {
		// y is line index
		if len(grid) == 0 {
			grid = make([][]int, len(line))
		}
		for x, char := range line {
			// x is char index
			if grid[y] == nil {
				grid[y] = make([]int, len(lines))
			}
			if n, err := strconv.Atoi(string(char)); err == nil {
				grid[y][x] = n
			} else {
				log.Fatal(err)
			}
		}
	}
	return grid
}

// grid[y] = make([]int, len(line))
// 		for x, char := range line {
// 			if n, err := strconv.Atoi(string(char); err != nil {
// 				grid[y][x] = n
// 			} else {
// 				log.Fatal(err)
// 			}

// 		}

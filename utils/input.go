package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadFileAsLines reads a file line by line and returns a slice of strings
func ReadFileAsLines(fileName string) []string {
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
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

// ReadFileAsRuneMatrix reads a file and returns a matrix of strings
// Returns as [x][y] where x is the column (horizontal) and y is the row (vertical)
func ReadFileAsRuneMatrix(fileName string) [][]rune {
	lines := ReadFileAsLines(fileName)
	var matrix [][]rune
	for _, line := range lines {
		// append each character in the line to the matrix
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	// convert matrix from [y][x] to [x][y]
	newMatrix := [][]rune{}
	for x := 0; x < len(matrix[0]); x++ {
		row := []rune{}
		for y := 0; y < len(matrix); y++ {
			row = append(row, matrix[y][x])
		}
		newMatrix = append(newMatrix, row)
	}
	return newMatrix
}

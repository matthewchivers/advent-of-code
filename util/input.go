package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

// ReadFileAsByteChunks reads a file and returns a slice of byte slices
func ReadFileAsByteChunks(fileName string) [][]byte {
	data := ReadFileAsBytes(fileName)
	var chunks [][]byte
	var chunk []byte
	for i := 0; i < len(data); i++ {
		if i == len(data)-1 || (data[i] == '\n' && data[i+1] == '\n') {
			if data[i] != '\n' {
				chunk = append(chunk, data[i])
			}
			chunks = append(chunks, chunk)
			chunk = []byte{}
			i++ // skip the second newline
			continue
		}
		chunk = append(chunk, data[i])
	}
	return chunks
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

func ParseIntList(line string, sep string) ([]int, error) {
	var nums []int
	for _, num := range strings.Split(line, sep) {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

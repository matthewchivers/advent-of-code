package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadFileAsString reads a file line by line and returns a slice of strings
func ReadFileAsString(fileName string) []string {
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

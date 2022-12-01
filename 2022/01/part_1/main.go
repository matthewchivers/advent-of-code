package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("../input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	maxCalories := 0
	highestElf := 0
	currentElf := 0
	currentCalories := 0
	for _, line := range lines {
		if line != "" {
			cals, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("problem converting string to int", err)
				os.Exit(1)
			}
			currentCalories += cals
		} else {
			currentElf++
			if currentCalories > maxCalories {
				maxCalories = currentCalories
				highestElf = currentElf
			}
			currentCalories = 0
		}
	}
	log.Printf("Elf %d has the highest calories with %d", highestElf, maxCalories)
}

func readLines(fileName string) ([]string, error) {
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
		return nil, err
	}
	return lines, nil
}

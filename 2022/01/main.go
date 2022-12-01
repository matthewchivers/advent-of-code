package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	lines := readLines("input.txt")
	var calorieTotals *[]int = new([]int)
	currentCalories := 0
	for _, line := range lines {
		if line != "" {
			cals := stringToInt(line)
			currentCalories += cals
		} else {
			*calorieTotals = append(*calorieTotals, currentCalories)
			currentCalories = 0
		}
	}
	sort.Ints(*calorieTotals)
	topThreeTotal := 0
	for i := 0; i < 3; i++ {
		lastVal := pop(calorieTotals)
		topThreeTotal += lastVal
	}

	log.Println("Top Three Total: ", topThreeTotal)
}

func stringToInt(line string) int {
	val, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal("problem converting string to int", err)
	}
	return val
}

func pop(s *[]int) int {
	backVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return backVal
}

func readLines(fileName string) []string {
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

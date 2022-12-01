package main

import (
	"log"
	"sort"

	. "github.com/matthewchivers/advent-of-code/lib-go"
)

func main() {
	lines := ReadLines("../input.txt")
	var calorieTotals *[]int = new([]int)
	currentCalories := 0
	for _, line := range lines {
		if line != "" {
			cals := StringToInt(line)
			currentCalories += cals
		} else {
			*calorieTotals = append(*calorieTotals, currentCalories)
			currentCalories = 0
		}
	}
	sort.Ints(*calorieTotals)
	topThreeTotal := 0
	for i := 0; i < 3; i++ {
		lastVal := SlicePopInt(calorieTotals)
		topThreeTotal += lastVal
	}

	log.Println("Top Three Total: ", topThreeTotal)
}

package main

import (
	"fmt"

	"github.com/matthewchivers/advent-of-code/utils"
)

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 1!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two: Not yet implemented")
}

func partOne() int {
	lines := utils.ReadFileAsString("input.txt")
	for _, line := range lines {
		fmt.Println(line)
	}
	return 0
}

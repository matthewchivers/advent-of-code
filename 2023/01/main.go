package main

import (
	"fmt"

	"github.com/matthewchivers/advent-of-code/utils"
)

func main() {
	fmt.Println("Hello, advent of code 2023!")
	lines := utils.ReadFileAsString("input.txt")
	for _, line := range lines {
		fmt.Println(line)
	}
}

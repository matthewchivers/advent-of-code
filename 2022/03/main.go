package main

import (
	"log"
	"strings"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

func main() {
	lines := aoc.ReadLines("input.txt")
	total := 0
	for _, line := range lines {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]
		for _, char := range secondHalf {
			if strings.Contains(firstHalf, string(char)) {
				if char > 96 {
					total += (int(char) - 96)
				} else {
					total += (int(char) - 38)
				}
				break
			}
		}
	}
	log.Println("Total:", total)
}

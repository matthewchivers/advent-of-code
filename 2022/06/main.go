package main

import (
	"fmt"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

var (
	data = aoc.ReadFileAsBytes("input.txt")
)

func main() {
	fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo())
}

func partOne() int {
	return getFirstUnique(4)
}

func partTwo() int {
	return getFirstUnique(14)
}

// getFirstUnique returns the index of the first window of size `size` that contains no duplicate bytes
// If no such window exists, -1 is returned
// The function uses a sliding window approach to keep track of the frequency of bytes in the window
// and the number of duplicate bytes in the window
func getFirstUnique(size int) int {
	var freq [256]int // count occurrences of each byte (0-256) in the window
	duplicates := 0   // track number of duplicate bytes in the window

	for i := 0; i < len(data); i++ {
		b := data[i]
		freq[b]++ // add byte to window

		if freq[b] == 2 {
			// byte addition has caused a duplicate - increment the count
			duplicates++
		}

		// if the window is now larger than the required size, remove the byte at the start of the window
		if i >= size {
			prevByte := data[i-size]
			freq[prevByte]--

			// byte removal has caused the byte to no longer be a duplicate - decrement the count
			if freq[prevByte] == 1 {
				duplicates--
			}
		}

		if i >= size-1 && duplicates == 0 {
			return i + 1
		}
	}
	return -1
}

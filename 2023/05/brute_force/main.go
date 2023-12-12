package main

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/advent-of-code/utils"
)

// Mapping is a struct to hold the mapping data
type Mapping struct {
	destination int
	source      int
	offset      int
}

var mappings [][]Mapping

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 5!")
	fmt.Println("Part one:", partOne())
	// fmt.Println("Part two:", partTwo())
}

// PART ONE WORKS BUT IT ISN'T SCALABLE (AS SHOWN BY PART TWO)
// partOne returns the answer to part one of the day's puzzle
func partOne() int {
	lowestLocation := -1

	// chunks := utils.ReadFileAsByteChunks("exInput.txt")
	chunks := utils.ReadFileAsByteChunks("input.txt")

	seeds, err := utils.StringToIntArray(strings.SplitN(string(chunks[0]), ": ", 2)[1])
	if err != nil {
		panic(err)
	}

	for i := 1; i < len(chunks); i++ {
		chunkMap := []Mapping{}
		for _, line := range strings.Split(string(chunks[i]), "\n")[1:] {
			nums, err := utils.StringToIntArray(line)
			if err != nil {
				panic(err)
			}
			chunkMap = append(chunkMap, Mapping{
				destination: nums[0],
				source:      nums[1],
				offset:      nums[2],
			})
		}
		mappings = append(mappings, chunkMap)
	}

	for _, seed := range seeds {
		transformHistory := fmt.Sprintf("%d", seed)
		transformedSeed := seed
		fmt.Printf("Seed before: %d\n", transformedSeed)
		for _, mapping := range mappings {
			for _, mapItem := range mapping {
				if transformedSeed >= mapItem.source && transformedSeed <= mapItem.source+(mapItem.offset-1) {
					fmt.Printf("Mapping: %d -> %d, offset %d\n", mapItem.source, mapItem.destination, mapItem.offset)
					before := transformedSeed
					transformedSeed = mapItem.destination + (transformedSeed - mapItem.source)
					after := transformedSeed
					fmt.Printf("Transformed: %d -> %d\n\n", before, after)
					break
				}
			}
			transformHistory = fmt.Sprintf("%s, %d", transformHistory, transformedSeed)
		}
		fmt.Printf("Seed after: %d, History: %s\n", transformedSeed, transformHistory)
		if lowestLocation == -1 || transformedSeed < lowestLocation {
			lowestLocation = transformedSeed
		}

	}
	return lowestLocation
}

// partTwo returns the answer to part two of the day's puzzle
func partTwo() int {
	lowestLocation := -1

	// chunks := utils.ReadFileAsByteChunks("exInput.txt")
	chunks := utils.ReadFileAsByteChunks("input.txt")

	seedsList, err := utils.StringToIntArray(strings.SplitN(string(chunks[0]), ": ", 2)[1])
	if err != nil {
		panic(err)
	}
	for i := 1; i < len(chunks); i++ {
		chunkMap := []Mapping{}
		for _, line := range strings.Split(string(chunks[i]), "\n")[1:] {
			nums, err := utils.StringToIntArray(line)
			if err != nil {
				panic(err)
			}
			chunkMap = append(chunkMap, Mapping{
				destination: nums[0],
				source:      nums[1],
				offset:      nums[2],
			})
		}
		mappings = append(mappings, chunkMap)
	}

	for i := 0; i < len(seedsList); i += 2 {
		for j := 0; j < seedsList[i+1]; j++ {
			seed := seedsList[i] + j

			// for _, seed := range seeds {
			transformHistory := fmt.Sprintf("%d", seed)
			transformedSeed := seed
			fmt.Printf("Seed before: %d\n", transformedSeed)
			for _, mapping := range mappings {
				for _, mapItem := range mapping {
					if transformedSeed >= mapItem.source && transformedSeed <= mapItem.source+(mapItem.offset-1) {
						fmt.Printf("Mapping: %d -> %d, offset %d\n", mapItem.source, mapItem.destination, mapItem.offset)
						before := transformedSeed
						transformedSeed = mapItem.destination + (transformedSeed - mapItem.source)
						after := transformedSeed
						fmt.Printf("Transformed: %d -> %d\n\n", before, after)
						break
					}
				}
				transformHistory = fmt.Sprintf("%s, %d", transformHistory, transformedSeed)
			}
			fmt.Printf("Seed after: %d, History: %s\n", transformedSeed, transformHistory)
			if lowestLocation == -1 || transformedSeed < lowestLocation {
				lowestLocation = transformedSeed
			}

		}
	}
	return lowestLocation

}

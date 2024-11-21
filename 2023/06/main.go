package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/matthewchivers/advent-of-code/util"
)

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 6!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

type Race struct{ Time, Distance int }

// Part One returns the answer to part one of the day's puzzle
func partOne() int {
	lines := util.ReadFileAsLines("input.txt")
	races, err := parseRaces(lines[0], lines[1], false)
	if err != nil {
		panic(err)
	}
	totalWays := 1
	for _, race := range races {
		totalWays *= calculateWins(race)
	}
	return totalWays
}

// Part Two returns the answer to part two of the day's puzzle
func partTwo() int {
	lines := util.ReadFileAsLines("input.txt")
	races, err := parseRaces(lines[0], lines[1], true)
	if err != nil {
		panic(err)
	}
	race := races[0]
	return calculateWins(race)
}

func calculateWins(race Race) int {
	// return quadraticWins(race)
	return bruteForceWins(race)
}

func quadraticWins(race Race) int {
	// Quadratic formula : t = (-b +- sqrt(b^2 - 4ac)) / 2a
	a := -1
	b := race.Time
	c := -race.Distance

	discriminant := b*b - 4*a*c

	// Check for non-negative discriminant
	if discriminant < 0 {
		return 0
	}

	// Calculating roots
	sqrtDisc := math.Sqrt(float64(discriminant))
	root1 := (float64(-b) + sqrtDisc) / (2 * float64(a))
	root2 := (float64(-b) - sqrtDisc) / (2 * float64(a))

	// Ensure root1 is the smaller root
	if root1 > root2 {
		root1, root2 = root2, root1
	}

	// Count integer solutions within the range
	start := int(math.Ceil(root1))
	end := int(math.Floor(root2))
	count := 0
	for t := start; t <= end; t++ {
		if t >= 0 && t <= race.Time {
			count++
		}
	}
	return count
}

func bruteForceWins(race Race) int {
	// Bruteforce
	waysToWin := 0
	winTriggered := false

	for t := 0; t < race.Time; t++ {
		// distance travelled = t * (T - t) (where T is the time limit and t is the time the button is held down
		if distance := t * (race.Time - t); distance > race.Distance {
			waysToWin++
		} else if winTriggered {
			break
		}
	}
	return waysToWin
}

func parseRaces(times, distances string, stripSpaces bool) ([]Race, error) {
	timeString := strings.SplitN(times, ":", -1)[1]
	distanceString := strings.SplitN(distances, ":", -1)[1]
	if stripSpaces {
		timeString = strings.ReplaceAll(timeString, " ", "")
		distanceString = strings.ReplaceAll(distanceString, " ", "")
	}
	timeVals, err := util.StringToIntArray(timeString)
	if err != nil {
		return nil, err
	}
	distanceVals, err := util.StringToIntArray(distanceString)
	if err != nil {
		return nil, err
	}
	races := []Race{}
	for i := 0; i < len(timeVals); i++ {
		time := timeVals[i]
		distance := distanceVals[i]

		races = append(races, Race{
			Time:     time,
			Distance: distance,
		})
	}
	return races, nil
}

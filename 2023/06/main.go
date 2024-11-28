package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/matthewchivers/advent-of-code/util"
)

var (
	lines = util.ReadFileAsLines("input.txt")
)

func main() {
	fmt.Println("Hello, Advent of Code 2023 - Day 6!")

	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

type Race struct {
	Time, Distance int
}

// partOne calculates the total number of ways to win all races combined
func partOne() int {
	races, err := parseRaces(lines[0], lines[1], false)
	if err != nil {
		panic(err)
	}
	totalWays := 1
	for _, race := range races {
		totalWays *= calculateWinningWays(race)
		if totalWays == 0 {
			break // Exit early if no ways to win exist
		}
	}
	return totalWays
}

// partTwo calculates the number of ways to win a single combined race
func partTwo() int {
	races, err := parseRaces(lines[0], lines[1], true)
	if err != nil {
		panic(err)
	}
	return calculateWinningWays(races[0])
}

// calculateWinningWays calculates the number of ways to win a given race using a quadratic equation
func calculateWinningWays(race Race) int {
	return quadraticSolution(race)
}

// quadraticSolution solves the quadratic equation to determine the number of valid winning times
func quadraticSolution(race Race) int {
	// Quadratic formula: t = (-b ± sqrt(b^2 - 4ac)) / 2a
	// with t being the time to hold the button
	// t_1 (root 1) = (-b + sqrt(b^2 - 4ac)) / 2a
	// t_2 (root 2) = (-b - sqrt(b^2 - 4ac)) / 2a

	// Solving: t * (T - t) > Distance
	// Equivalent to: t^2 - tT - Distance > 0
	// t = time to hold the button
	//   also millimetres per millisecond speed (mm/ms)
	// T = total race time allowed
	// Distance = distance to beat

	// Coefficients of the quadratic equation
	a := 1             // Coefficient of t (where t^2 = 1t^2)
	b := -race.Time    // -T = Coefficient of t (-tT)
	c := race.Distance // Constant term (D = Distance to beat)

	// b^2 - 4ac
	discriminant := b*b - 4*a*c

	// No valid solutions if discriminant is negative
	if discriminant < 0 {
		return 0
	}

	// Calculate roots of the equation
	sqrtDisc := math.Sqrt(float64(discriminant))
	root1 := (float64(-b) + sqrtDisc) / (2 * float64(a))
	root2 := (float64(-b) - sqrtDisc) / (2 * float64(a))

	// Ensure root1 is the smaller root
	if root1 > root2 {
		root1, root2 = root2, root1
	}

	// Calculate the range of valid integer solutions
	start := int(math.Ceil(root1))
	end := int(math.Floor(root2))

	// Return 0 if no valid range exists
	if start > end || start > race.Time {
		return 0
	}

	return end - start + 1
}

// parseRaces converts input strings to Race structs
func parseRaces(times, distances string, stripSpaces bool) ([]Race, error) {
	timeString := strings.SplitN(times, ":", 2)[1]
	distanceString := strings.SplitN(distances, ":", 2)[1]
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

	races := make([]Race, len(timeVals))
	for i := range timeVals {
		races[i] = Race{
			Time:     timeVals[i],
			Distance: distanceVals[i],
		}
	}
	return races, nil
}

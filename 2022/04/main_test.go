package main

import (
	"testing"

	aoc "github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	input := aoc.ReadFileAsLines("sample.txt")
	expected := 2
	result := partOne(input)

	assert.Equal(t, expected, result, "solvePart(1) should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	input := aoc.ReadFileAsLines("sample.txt")
	expected := 4
	result := partTwo(input)

	assert.Equal(t, expected, result, "solvePart(2) should return the correct total value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := aoc.ReadFileAsLines("sample.txt")
		partOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := aoc.ReadFileAsLines("sample.txt")
		partTwo(input)
	}
}

package main

import (
	"testing"

	aoc "github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	input := aoc.ReadFileAsLines("sample.txt")
	expected := "CMZ"
	result := partOne(input)

	assert.Equal(t, expected, result, "partOne() should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	input := aoc.ReadFileAsLines("sample.txt")
	expected := "MCD"
	result := partTwo(input)

	assert.Equal(t, expected, result, "partTwo() should return the correct total value")
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
		partOne(input)
	}
}

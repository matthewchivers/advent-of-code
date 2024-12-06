package main

import (
	"testing"

	aoc "github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

var (
	input []string = aoc.ReadFileAsLines("sample.txt")
)

func TestCalculatePartOne(t *testing.T) {
	expected := 15
	result := processLines(1, input)

	assert.Equal(t, expected, result, "calculatePartOne should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	expected := 12
	result := processLines(2, input)

	assert.Equal(t, expected, result, "calculatePartTwo should return the correct total value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := aoc.ReadFileAsLines("sample.txt")
		processLines(1, input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := aoc.ReadFileAsLines("sample.txt")
		processLines(2, input)
	}
}

package main

import (
	"testing"

	aoc "github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

var ()

func TestCalculatePartOne(t *testing.T) {

	input := aoc.ReadFileAsLines("sample.txt")
	expected := 24000
	result := calculatePartOne(input)

	assert.Equal(t, expected, result, "calculatePartOne should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	input := aoc.ReadFileAsLines("sample.txt")
	expected := 45000
	result := calculatePartTwo(input)

	assert.Equal(t, expected, result, "calculatePartTwo should return the correct total value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := aoc.ReadFileAsLines("sample.txt")
		calculatePartOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := aoc.ReadFileAsLines("sample.txt")
		calculatePartTwo(input)
	}
}

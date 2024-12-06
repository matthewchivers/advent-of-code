package main

import (
	"testing"

	aoc "github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	lines := aoc.ReadFileAsLines("sample.txt")
	expected := 142
	result := partOne(lines)

	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestCalculatePartTwo(t *testing.T) {
	lines := aoc.ReadFileAsLines("sample2.txt")
	expected := 281
	result := partTwo(lines)

	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := aoc.ReadFileAsLines("sample.txt")
		partOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := aoc.ReadFileAsLines("sample2.txt")
		partTwo(lines)
	}
}

package main

import (
	"testing"

	aoc "github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPartOneSample(t *testing.T) {
	lines := aoc.ReadFileAsLines("sample.txt")
	expected := 11
	result := partOne(lines)
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwoSample(t *testing.T) {
	lines := aoc.ReadFileAsLines("sample.txt")
	expected := 31
	result := partTwo(lines)
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func TestPartOneInput(t *testing.T) {
	lines := aoc.ReadFileAsLines("input.txt")
	expected := 1506483
	result := partOne(lines)
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwoInput(t *testing.T) {
	lines := aoc.ReadFileAsLines("input.txt")
	expected := 23126924
	result := partTwo(lines)
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := aoc.ReadFileAsLines("input.txt")
		partOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := aoc.ReadFileAsLines("input.txt")
		partTwo(lines)
	}
}

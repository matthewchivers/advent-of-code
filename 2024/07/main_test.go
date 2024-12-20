package main

import (
	"testing"

	"github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPartOneSample(t *testing.T) {
	lines := util.ReadFileAsLines("sample.txt")
	expected := 3749
	result := partOne(lines)
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwoSample(t *testing.T) {
	lines := util.ReadFileAsLines("sample.txt")
	expected := 11387
	result := partTwo(lines)
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func TestPartOneInput(t *testing.T) {
	lines := util.ReadFileAsLines("input.txt")
	expected := 3245122495150
	result := partOne(lines)
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwoInput(t *testing.T) {
	lines := util.ReadFileAsLines("input.txt")
	expected := 105517128211543
	result := partTwo(lines)
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := util.ReadFileAsLines("input.txt")
		partOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := util.ReadFileAsLines("input.txt")
		partTwo(lines)
	}
}

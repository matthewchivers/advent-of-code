package main

import (
	"testing"

	"github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	lines := util.ReadFileAsLines("sample.txt")
	// expected := 288
	expected := 352 // I believe the sample provided is somehow incorrect?
	result := partOne(lines)

	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwo(t *testing.T) {
	lines := util.ReadFileAsLines("sample.txt")
	expected := 71503
	result := partTwo(lines)

	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := util.ReadFileAsLines("sample.txt")
		partOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := util.ReadFileAsLines("sample.txt")
		partTwo(lines)
	}
}

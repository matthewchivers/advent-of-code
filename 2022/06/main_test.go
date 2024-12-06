package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

func TestCalculatePartOne(t *testing.T) {
	data := aoc.ReadFileAsBytes("sample.txt")
	expected := 11
	result := partOne(data)

	assert.Equal(t, expected, result, "partOne() should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	data := aoc.ReadFileAsBytes("sample.txt")
	expected := 26
	result := partTwo(data)

	assert.Equal(t, expected, result, "partTwo() should return the correct total value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := aoc.ReadFileAsBytes("sample.txt")
		partOne(data)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := aoc.ReadFileAsBytes("sample.txt")
		partOne(data)
	}
}

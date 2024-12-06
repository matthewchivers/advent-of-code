package main

import (
	"testing"

	"github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	matrix := util.ReadFileAsRuneMatrix("sample.txt")
	expected := 4361
	result := partOne(matrix)

	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestCalculatePartTwo(t *testing.T) {
	matrix := util.ReadFileAsRuneMatrix("sample.txt")
	expected := 467835
	result := partTwo(matrix)

	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matrix := util.ReadFileAsRuneMatrix("sample.txt")
		partOne(matrix)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matrix := util.ReadFileAsRuneMatrix("sample.txt")
		partTwo(matrix)
	}
}

package main

import (
	"testing"

	"github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	lines := util.ReadFileAsLines("sample.txt")
	expected := 2
	result := partOne(lines)
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwo(t *testing.T) {
	lines := util.ReadFileAsLines("sample.txt")
	expected := 4
	result := partTwo(lines)
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	lines := util.ReadFileAsLines("sample.txt")
	for n := 0; n < b.N; n++ {
		partOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	lines := util.ReadFileAsLines("sample.txt")
	for n := 0; n < b.N; n++ {
		partTwo(lines)
	}
}

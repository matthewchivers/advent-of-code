package main

import (
	"testing"

	"github.com/matthewchivers/advent-of-code/util"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	chunks := util.ReadFileAsByteChunks("sample.txt")
	expected := 35
	result := partOne(chunks)

	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwo(t *testing.T) {
	chunks := util.ReadFileAsByteChunks("sample.txt")
	expected := 46
	result := partTwo(chunks)

	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		chunks := util.ReadFileAsByteChunks("sample.txt")
		partOne(chunks)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		chunks := util.ReadFileAsByteChunks("sample.txt")
		partTwo(chunks)
	}
}

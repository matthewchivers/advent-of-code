package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	expected := 639
	result := partOne()
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwo(t *testing.T) {
	expected := 674
	result := partTwo()
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		partOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		partTwo()
	}
}

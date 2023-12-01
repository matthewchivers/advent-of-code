package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	expected := 55488
	result := partOne()

	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestCalculatePartTwo(t *testing.T) {
	expected := 55614
	result := partTwo()

	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

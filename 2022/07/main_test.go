package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	expected := 1989474
	result := partOne()

	assert.Equal(t, expected, result, "partOne() should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	expected := 1111607
	result := partTwo()

	assert.Equal(t, expected, result, "partTwo() should return the correct total value")
}

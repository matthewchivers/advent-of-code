package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	expected := "RLFNRTNFB"
	result := partOne()

	assert.Equal(t, expected, result, "partOne() should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	expected := "MHQTLJRLB"
	result := partTwo()

	assert.Equal(t, expected, result, "partTwo() should return the correct total value")
}

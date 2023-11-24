package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	expected := 515
	result := solvePart(1)

	assert.Equal(t, expected, result, "solvePart(1) should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	expected := 883
	result := solvePart(2)

	assert.Equal(t, expected, result, "solvePart(2) should return the correct total value")
}

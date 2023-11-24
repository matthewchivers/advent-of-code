package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	expected := 9759
	result := processLines(1)

	assert.Equal(t, expected, result, "calculatePartOne should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	expected := 12429
	result := processLines(2)

	assert.Equal(t, expected, result, "calculatePartTwo should return the correct total value")
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	// Reset the podium and populated flag before each test
	podium = Podium{}
	populated = false
	lines = nil // Ensure lines are loaded from the file

	expected := 71780
	result := calculatePartOne()

	assert.Equal(t, expected, result, "calculatePartOne should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	// Reset the podium and populated flag before each test
	podium = Podium{}
	populated = false
	lines = nil // Ensure lines are loaded from the file

	expected := 212489
	result := calculatePartTwo()

	assert.Equal(t, expected, result, "calculatePartTwo should return the correct total value")
}

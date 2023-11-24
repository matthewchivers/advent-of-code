package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePartOne(t *testing.T) {
	expected := 71780
	result := calculatePartOne()

	assert.Equal(t, expected, result, "calculatePartOne should return the correct highest value")
}

func TestCalculatePartTwo(t *testing.T) {
	expected := 212489
	result := calculatePartTwo()

	assert.Equal(t, expected, result, "calculatePartTwo should return the correct total value")
}

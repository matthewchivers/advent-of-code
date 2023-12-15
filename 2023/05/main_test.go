package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	expected := 3374647
	result := partOne()

	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwo(t *testing.T) {
	expected := 6082852
	result := partTwo()

	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

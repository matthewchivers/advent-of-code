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

// Disabled part two because it takes far too long to run

// func TestPartTwo(t *testing.T) {
// 	expected := 0
// 	result := partTwo()

// 	assert.Equal(t, expected, result, "partTwo() should return the correct value")
// }

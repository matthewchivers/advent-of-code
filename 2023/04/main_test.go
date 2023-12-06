package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	expected := 21919
	result := partOne()

	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwo(t *testing.T) {
	expected := 9881048
	result := partTwo()

	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

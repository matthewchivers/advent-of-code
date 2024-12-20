package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneSample(t *testing.T) {
	f, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	expected := 161
	result := partOne(r)
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwoSample(t *testing.T) {
	f, err := os.Open("sample2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	expected := 48
	result := partTwo(r)
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func TestPartOneInput(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	expected := 173529487
	result := partOne(r)
	assert.Equal(t, expected, result, "partOne() should return the correct value")
}

func TestPartTwoInput(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	expected := 99532691
	result := partTwo(r)
	assert.Equal(t, expected, result, "partTwo() should return the correct value")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r := bufio.NewReader(f)
		partOne(r)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r := bufio.NewReader(f)
		partTwo(r)
	}
}

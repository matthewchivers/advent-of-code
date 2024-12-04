package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

func main() {
	fmt.Println("Hello, advent of code 2024 - Day 3!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

func partOne() int {
	return processInput(true)
}

func partTwo() int {
	return processInput(false)
}

func processInput(isPartOne bool) int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var buffer strings.Builder

	// define valid commands
	mulCommand := []byte("mul")
	doCommand := []byte("do")
	dontCommand := []byte("don't")

	num1 := 0  // first number for multiplication
	num2 := 0  // second number for multiplication
	total := 0 // accumulator for the sum of products

	multiply := true // keeps track of whether `mul` commands are active

	// state machine:
	// 0 - searching for command letters and opening bracket
	// 1 - reading numbers between brackets for multiplication
	// 2 - processing `do()` command - turns on multiplication
	// 3 - processing `don't()` command - turns off multiplication
	state := 0 // initial state for command parsing

	for {
		switch state {
		case 0:
			char, err := reader.ReadByte()
			if err != nil {
				return total // end of file
			}
			// detect letters of a valid command (`mul`, `do`, `don't`)
			if (buffer.Len() < len(mulCommand) && char == mulCommand[buffer.Len()]) ||
				(buffer.Len() < len(doCommand) && char == doCommand[buffer.Len()]) ||
				(buffer.Len() < len(dontCommand) && char == dontCommand[buffer.Len()]) {
				buffer.WriteByte(char)
				continue
			}
			// if opening bracket is detected, check command and transition state
			if char == '(' {
				if buffer.String() == string(mulCommand) {
					// only proceed with multiplication if it's enabled
					if multiply {
						state = 1
						buffer.Reset()
						continue
					}
				}
				if isPartOne {
					// just do part one
					state = 0
					buffer.Reset()
					continue
				}
				// handle `do` and `don't` commands
				if buffer.String() == string(doCommand) {
					state = 2
				} else if buffer.String() == string(dontCommand) {
					state = 3
				}
			}
			// reset the buffer before transitioning to the next command
			// or if unexpected character is found
			buffer.Reset()
		case 1: // `mul` - read numbers between brackets
			char, err := reader.ReadByte()
			if err != nil {
				return total // end of file
			}
			// collect digits for the number
			if char >= '0' && char <= '9' {
				buffer.WriteByte(char)
				continue
			}
			// if comma is found, store the first number
			if char == ',' {
				num, err := aoc.StringToInt(buffer.String())
				if err != nil {
					panic(err)
				}
				num1 = num
				buffer.Reset()
				continue
			}
			// if closing bracket is found, complete the multiplication
			if char == ')' {
				num, err := aoc.StringToInt(buffer.String())
				if err != nil {
					panic(err)
				}
				num2 = num
				// add the product of the two numbers to the total
				total += num1 * num2
				state = 0 // reset to search for next command
				buffer.Reset()
				continue
			}
			// if unexpected character is found, reset state and buffer
			buffer.Reset()
			state = 0
		case 2: // `do()` - enable multiplication
			char, err := reader.ReadByte()
			if err != nil {
				return total // end of file
			}
			if char == ')' {
				multiply = true // enable multiplication for future `mul` commands
				state = 0
				continue
			}
			state = 0 // unexpected character, reset state
		case 3: // `don't()` - disable multiplication
			char, err := reader.ReadByte()
			if err != nil {
				return total // end of file
			}
			if char == ')' {
				multiply = false // disable multiplication for future `mul` commands
				state = 0
				continue
			}
			state = 0 // unexpected character, reset state
		}
	}
}

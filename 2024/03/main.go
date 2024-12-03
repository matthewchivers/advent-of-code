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

	mulCommand := []byte("mul")
	doCommand := []byte("do")
	dontCommand := []byte("don't")

	state := 0
	num1 := 0
	num2 := 0
	total := 0

	multiply := true

	// state machine:
	// 0 - searching for command letter(s) and brackets start
	// 1 - reading between brackets (valid chars are numbers, comma, and end-bracket)
	// 2 - do command - turn on
	// 3 - don't command - turn off
	// 4 - performing multiplication

	for {
		switch state {
		case 0:
			char, err := reader.ReadByte()
			if err != nil {
				return total // End of file
			}
			if (buffer.Len() < len(mulCommand) && char == mulCommand[buffer.Len()]) ||
				(buffer.Len() < len(doCommand) && char == doCommand[buffer.Len()]) ||
				(buffer.Len() < len(dontCommand) && char == dontCommand[buffer.Len()]) {
				buffer.WriteByte(char)
				continue
			}
			if char == '(' {
				if buffer.String() == string(mulCommand) {
					if multiply {
						state = 1
						buffer.Reset()
						continue
					}
				}
				if isPartOne {
					state = 0
					buffer.Reset()
					continue
				}
				if buffer.String() == string(doCommand) {
					state = 2
				} else if buffer.String() == string(dontCommand) {
					state = 3
				}
			}
			buffer.Reset()
		case 1: // mul - get everything in brackets
			char, err := reader.ReadByte()
			if err != nil {
				return total // End of file
			}
			if char >= '0' && char <= '9' {
				buffer.WriteByte(char)
				continue
			}
			if char == ',' {
				num, err := aoc.StringToInt(buffer.String())
				if err != nil {
					panic(err)
				}
				num1 = num
				buffer.Reset()
				continue
			}
			if char == ')' {
				num, err := aoc.StringToInt(buffer.String())
				if err != nil {
					panic(err)
				}
				num2 = num
				total += num1 * num2
				state = 0
				buffer.Reset()
				continue
			}
			buffer.Reset()
			state = 0
		case 2: // do - search for end bracket and turn on
			char, err := reader.ReadByte()
			if err != nil {
				return total // End of file
			}
			if char == ')' {
				multiply = true
				state = 0
				continue
			}
			state = 0
		case 3: // don't - search for end bracket and turn off
			char, err := reader.ReadByte()
			if err != nil {
				return total // End of file
			}
			if char == ')' {
				multiply = false
				state = 0
				continue
			}
			state = 0
		}
	}
}

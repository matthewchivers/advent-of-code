package main

import (
	"log"
	"regexp"
	"strconv"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

type instruction struct {
	count int
	from  int
	to    int
}

type status int

const (
	stackInput       status = iota // Represents stack input parsing phase
	instructionInput               // Represents instruction parsing phase
)

func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	log.Println("Part One: ", partOne(lines))
	log.Println("Part Two: ", partTwo(lines))
}

// partOne simulates moving crates one at a time based on the parsed instructions
func partOne(input []string) string {
	stacks, instructions := parseLines(input)
	for _, inst := range instructions {
		moveCrates(stacks, inst, false)
	}
	return getTopBoxes(stacks)
}

// partTwo simulates moving multiple crates at once using a buffer stack
func partTwo(input []string) string {
	stacks, instructions := parseLines(input)
	for _, inst := range instructions {
		moveCrates(stacks, inst, true)
	}
	return getTopBoxes(stacks)
}

// moveCrates moves crates from one stack to another
// If useBuffer is true, crates are moved using a temporary buffer to simulate moving multiple crates at once.
func moveCrates(stacks []aoc.Stack[rune], inst instruction, useBuffer bool) {
	if useBuffer {
		// Use a buffer stack to hold crates temporarily before moving to the destination
		// e.g. From[ A, B, C, D ] => Buff[ D, C, B, A] = To[ A, B, C, D]
		buffer := aoc.Stack[rune]{}
		for i := 0; i < inst.count; i++ {
			buffer.Push(stacks[inst.from-1].Pop()) // Pop crates from the source stack into the buffer
		}
		for i := 0; i < inst.count; i++ {
			stacks[inst.to-1].Push(buffer.Pop()) // Push crates from the buffer to the destination stack
		}
	} else {
		// Move crates directly from the source stack to the destination stack, one at a time
		// e.g. From[ A, B, C, D ] => To[ D, C, B, A ]
		for i := 0; i < inst.count; i++ {
			stacks[inst.to-1].Push(stacks[inst.from-1].Pop())
		}
	}
}

// getTopBoxes returns a string representing the top element of each stack
func getTopBoxes(outputStack []aoc.Stack[rune]) string {
	var output string
	for _, stack := range outputStack {
		if !stack.IsEmpty() {
			output += string(stack.Peek())
		}
	}
	return output
}

// parseLines parses the input lines to create stacks and instruction sets
func parseLines(input []string) ([]aoc.Stack[rune], []instruction) {
	stacks := make([]aoc.Stack[rune], (len(input[0])/4 + 1))
	instructions := []instruction{}
	currentStatus := stackInput
	for _, line := range input {
		if line != "" {
			switch currentStatus {
			case stackInput:
				currentStatus = parseLineStack(line, &stacks)
			case instructionInput:
				instructions = append(instructions, parseLineInstruction(line))
			}
		}
	}
	// Reverse the order of the stacks to match expected final order
	for _, stack := range stacks {
		stack.Reverse()
	}
	return stacks, instructions
}

// parseLineStack parses a line containing stack information and adds elements to stacks
func parseLineStack(line string, stacks *[]aoc.Stack[rune]) status {
	for i, character := range line {
		if (i-1)%4 == 0 {
			if character == '1' {
				return instructionInput
			}
			if character != ' ' {
				(*stacks)[i/4].Push(character)
			}
		}
	}
	return stackInput
}

// parseLineInstruction parses a line containing an instruction and returns an instruction struct
func parseLineInstruction(line string) instruction {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)
	count, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])

	return instruction{count, from, to}
}

package main

import (
	"log"
	"regexp"
	"strconv"

	aoc "github.com/matthewchivers/advent-of-code/utils"
)

type instruction struct {
	count int
	from  int
	to    int
}

type status int

const (
	stackInput       status = iota
	instructionInput status = iota
)

var (
	lines = aoc.ReadFileAsString("input.txt")
)

func main() {
	log.Println("Part One: ", partOne())
	log.Println("Part Two: ", partTwo())
}

func partOne() string {
	// move one crate at a time
	stacks, instructions := parseLines()
	for _, instruction := range instructions {
		for i := 0; i < instruction.count; i++ {
			stacks[instruction.to-1].Push(stacks[instruction.from-1].Pop())
		}
	}
	output := getTopBoxes(stacks)
	return output
}

func partTwo() string {
	// move multiple crates at a time
	// (uses a buffer to replicate this kind of movement)
	stacks, instructions := parseLines()
	for _, instruction := range instructions {
		buffer := aoc.Stack[rune]{}
		for i := 0; i < instruction.count; i++ {
			buffer.Push(stacks[instruction.from-1].Pop())
		}
		for i := 0; i < instruction.count; i++ {
			stacks[instruction.to-1].Push(buffer.Pop())
		}
	}
	output := getTopBoxes(stacks)
	return output
}

func getTopBoxes(outputStack []aoc.Stack[rune]) string {
	output := ""
	for _, stack := range outputStack {
		output += string(stack.Peek())
	}
	return output
}

func parseLines() ([]aoc.Stack[rune], []instruction) {
	stacks := make([]aoc.Stack[rune], (len(lines[0])/4 + 1))
	instructions := []instruction{}
	currentStatus := stackInput
	for _, line := range lines {
		if line != "" {
			switch currentStatus {
			case stackInput:
				currentStatus = parseLineStack(line, &stacks)
			case instructionInput:
				instructions = append(instructions, parseLineInstruction(line))
			}
		}
	}
	// reverse the order of the stacks
	for _, stack := range stacks {
		stack.Reverse()
	}
	return stacks, instructions
}

// Example Line: [W] [N] [H]
// As long as we start at index 2 (W), every fourth character after that is a letter
func parseLineStack(line string, stacks *[]aoc.Stack[rune]) status {
	for i, character := range line {
		if (i-1)%4 == 0 { // every fourth element, starting with the second
			if character == '1' {
				return instructionInput
			}
			if character != ' ' {
				// Reading from top down, so, push new items to the bottom of the stack
				(*stacks)[i/4].Push(character)
			}
		}
	}
	return stackInput
}

func parseLineInstruction(line string) instruction {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)
	num1, _ := strconv.Atoi(matches[0])
	num2, _ := strconv.Atoi(matches[1])
	num3, _ := strconv.Atoi(matches[2])

	return instruction{num1, num2, num3}
}

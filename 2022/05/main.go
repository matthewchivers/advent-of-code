package main

import (
	"log"
	"regexp"
	"strconv"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
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
	lines = aoc.ReadLinesAsString("input.txt")
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
		buffer := aoc.RuneStack{}
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

func getTopBoxes(outputStack []aoc.RuneStack) string {
	output := ""
	for _, stack := range outputStack {
		output += string(stack.Peek())
	}
	return output
}

func parseLines() ([]aoc.RuneStack, []instruction) {
	stacks := make([]aoc.RuneStack, (len(lines[0])/4 + 1))
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
	return stacks, instructions
}

func parseLineStack(line string, stacks *[]aoc.RuneStack) status {
	for i, character := range line {
		if (i-1)%4 == 0 { // every fourth element, starting with the second
			if character == '1' {
				return instructionInput
			}
			if character != ' ' {
				(*stacks)[i/4].LIFOPush(character)
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

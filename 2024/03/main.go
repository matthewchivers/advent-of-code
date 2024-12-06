package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type parseState int

// Define the different states of the state machine used for parsing
const (
	stateScan          parseState = iota // Initial state, scanning for instructions
	stateM                               // Found 'm', possibly start of 'mul'
	stateMu                              // Found 'mu', possibly start of 'mul'
	stateMul                             // Found 'mul', expecting '(' next
	stateMulOpenParen                    // Found 'mul(', expecting first number
	stateMulNumber1                      // Parsing the first number of 'mul(x, y)'
	stateMulComma                        // Found ',', expecting second number
	stateMulNumber2                      // Parsing the second number of 'mul(x, y)'
	stateD                               // Found 'd', possibly start of 'do' or 'don't'
	stateDo                              // Found 'do', expecting '(' next
	stateDoOpenParen                     // Found 'do(', enabling future multiplications
	stateDon                             // Found 'don', possibly start of 'don't'
	stateDont                            // Found 'don't', expecting '(' next
	stateDontOpenParen                   // Found 'don't(', disabling future multiplications
)

const maxNumLength = 3 // Maximum length for numbers being parsed

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	fmt.Println("Part One:", partOne(r))
	fmt.Println("Part Two:", partTwo(r))
}

// partOne solves the first part of the puzzle by scanning the corrupted memory
// and summing up all valid `mul` instructions without considering any conditional statements.
func partOne(r *bufio.Reader) int {
	return solve(r, false)
}

// partTwo solves the second part of the puzzle by handling additional `do()` and `don't()`
// instructions, which enable or disable subsequent `mul` operations, respectively.
func partTwo(r *bufio.Reader) int {
	return solve(r, true)
}

func solve(r *bufio.Reader, handleDoDont bool) int {
	var (
		sum        int
		enabled    = true
		numBuffer  strings.Builder
		num1, num2 int
	)

	state := stateScan // Start in the scanning state

	for {
		rn, _, err := r.ReadRune()
		if err == io.EOF {
			break // Exit loop when end of file is reached
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("State: %v, Rune: %c\n", state, rn)

		switch state {
		case stateScan:
			switch rn {
			case 'm':
				// Found 'm', could be the start of 'mul'
				state = stateM
			case 'd':
				// Found 'd', could be the start of 'do' or 'don't'
				if handleDoDont {
					state = stateD
				}
			}
		case stateM:
			if rn == 'u' {
				// Found 'mu', could be the start of 'mul'
				state = stateMu
			} else {
				state = stateScan
			}
		case stateMu:
			if rn == 'l' {
				// Found 'mul', expecting '(' next
				state = stateMul
			} else {
				state = stateScan
			}
		case stateMul:
			if rn == '(' {
				// Found 'mul(', start parsing the first number
				numBuffer.Reset() // Clear the buffer for the new number
				state = stateMulOpenParen
			} else {
				state = stateScan
			}
		case stateMulOpenParen:
			if rn >= '0' && rn <= '9' {
				// Found a digit, start accumulating the first number
				numBuffer.WriteRune(rn)
				state = stateMulNumber1
			} else {
				state = stateScan
			}
		case stateMulNumber1:
			if rn >= '0' && rn <= '9' && numBuffer.Len() < maxNumLength {
				numBuffer.WriteRune(rn)
			} else if rn == ',' {
				// Finished parsing first number, convert it to an integer
				parsed, err := strconv.Atoi(numBuffer.String())
				if err != nil {
					panic(fmt.Sprintf("Failed to parse number: %v", err))
				}
				num1 = parsed
				numBuffer.Reset() // Clear the buffer for the second number
				state = stateMulComma
			} else {
				state = stateScan
			}
		case stateMulComma:
			if rn >= '0' && rn <= '9' {
				// Found a digit, start accumulating the second number
				numBuffer.WriteRune(rn)
				state = stateMulNumber2
			} else {
				state = stateScan
			}
		case stateMulNumber2:
			if rn >= '0' && rn <= '9' && numBuffer.Len() < maxNumLength {
				// Accumulate digits for the second number, up to max length
				numBuffer.WriteRune(rn)
			} else if rn == ')' {
				// Finished parsing second number, parse it to an integer
				parsed, err := strconv.Atoi(numBuffer.String())
				if err != nil {
					panic(fmt.Sprintf("Failed to parse number: %v", err))
				}
				num2 = parsed
				// If both numbers are valid and `mul` is enabled, add to sum
				if num1 >= 0 && num2 >= 0 && enabled {
					sum += num1 * num2
				}
				// Reset to scanning state
				state = stateScan
			} else {
				state = stateScan
			}
		case stateD:
			if rn == 'o' {
				// Found 'do', could be the start of 'do()'
				state = stateDo
			} else {
				state = stateScan
			}
		case stateDo:
			if rn == '(' {
				// Found 'do(', enable future multiplications
				state = stateDoOpenParen
			} else if rn == 'n' {
				// Found 'don', could be the start of 'don't'
				state = stateDon
			} else {
				state = stateScan
			}
		case stateDon:
			if rn == '\'' {
				// Found 'don'', could be the start of 'don't'
				state = stateDont
			} else {
				state = stateScan
			}
		case stateDont:
			if rn == 't' {
				// Found 'don't', expecting '(' next
				// Waiting for next character to determine state
			} else if rn == '(' {
				// Found 'don't(', disable future multiplications
				state = stateDontOpenParen
			} else {
				state = stateScan
			}
		case stateDontOpenParen:
			if rn == ')' {
				// Found 'don't()', disable `mul` operations
				enabled = false
				state = stateScan
			} else {
				state = stateScan
			}
		case stateDoOpenParen:
			if rn == ')' {
				// Found 'do()', enable `mul` operations
				enabled = true
				state = stateScan
			} else {
				state = stateScan
			}
		default:
			state = stateScan
		}
	}

	return sum
}

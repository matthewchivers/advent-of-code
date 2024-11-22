#!/bin/bash

# Usage: ./create_aoc.sh YEAR DAY [DIRECTORY]

# Check if at least two arguments are provided
if [ $# -lt 2 ]; then
    echo "Usage: $0 YEAR DAY [DIRECTORY]"
    exit 1
fi

YEAR=$1
DAY=$2
# strip leading zeros from DAY (sanitise user input)
DAY=$(echo $DAY | sed 's/^0*//')
DAYPADDED=$(printf "%02d" $DAY) # ironically, this will add leading zeros back in for usage in the directory name

# Base directory (default to current directory if not provided)
BASE_DIR=${3:-.}

# Full path to the target directory
TARGET_DIR="$BASE_DIR/$YEAR/$DAYPADDED"

# Create the directory if it doesn't exist
if [ -d "$TARGET_DIR" ]; then
    echo "Directory $TARGET_DIR already exists."
else
    mkdir -p "$TARGET_DIR"
    echo "Created directory $TARGET_DIR."
fi

# Function to create or append to a file if it doesn't exist
create_file() {
    local file_path=$1
    local content=$2
    if [ ! -f "$file_path" ]; then
        echo -e "$content" > "$file_path"
        echo "Created $file_path."
    else
        echo "$file_path already exists, leaving it unchanged."
    fi
}

# Content for main.go
MAIN_GO_CONTENT="package main

import (
    \"fmt\"
    aoc \"github.com/matthewchivers/advent-of-code/util\"
)

var (
	lines = aoc.ReadFileAsLines(\"input.txt\")
)

func main() {
    fmt.Println(\"Hello, advent of code $YEAR - Day $DAY!\")
	fmt.Println(\"Part one:\", partOne())
	fmt.Println(\"Part two:\", partTwo())
}

func partOne() int{
    fmt.Println(\"Part one not implemented\")
    return 0
}

func partTwo() int {
    fmt.Println(\"Part two not implemented\")
    return 0
}
"

# Content for main_test.go
MAIN_TEST_GO_CONTENT="package main

import (
	\"testing\"

	\"github.com/stretchr/testify/assert\"
)

func TestPartOne(t *testing.T) {
    t.Error(\"$YEAR Day $DAY TestPartOne not implemented\")
	
	expected := 0
    result := partOne()
	assert.Equal(t, expected, result, \"partOne() should return the correct value\")
}

func TestPartTwo(t *testing.T) {
	t.Error(\"$YEAR Day $DAY TestPartTwo not implemented\")

    expected := 0
    result := partTwo()
	assert.Equal(t, expected, result, \"partTwo() should return the correct value\")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		partOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		partTwo()
	}
}
"

# Content for README.md
README_CONTENT="# Advent of Code $YEAR - Day $DAY

[Link to the challenge](https://adventofcode.com/$YEAR/day/$DAY)

## Part One

## Part Two
"

# Create the files
create_file "$TARGET_DIR/main.go" "$MAIN_GO_CONTENT"
create_file "$TARGET_DIR/main_test.go" "$MAIN_TEST_GO_CONTENT"
create_file "$TARGET_DIR/input.txt" ""
create_file "$TARGET_DIR/README.md" "$README_CONTENT"

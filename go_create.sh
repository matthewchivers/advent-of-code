#!/bin/bash

# Usage: ./create_aoc.sh YEAR DAY [DIRECTORY]

# Check if at least two arguments are provided
if [ $# -lt 2 ]; then
    echo "Usage: $0 YEAR DAY [DIRECTORY]"
    exit 1
fi

YEAR=$1
DAY=$2
# Strip leading zeros from DAY (sanitise user input)
DAY=$(echo $DAY | sed 's/^0*//')
DAYPADDED=$(printf "%02d" $DAY) # Add leading zeros for usage in the directory name

# Base directory (default to current directory if not provided)
BASE_DIR=$(pwd)

# Full path to the target directory
TARGET_DIR="$BASE_DIR/$YEAR/$DAYPADDED"
INPUT_DIR="$BASE_DIR/input/$YEAR/$DAYPADDED"

create_target_dir() {
    if [ -d "$1" ]; then
        echo "Directory $1 already exists."
    else
        mkdir -p "$1"
        echo "Created directory $1."
    fi
}

create_target_dir "$TARGET_DIR"
create_target_dir "$INPUT_DIR"

# Function to create (not overwrite) a file from a template
create_from_template() {
    local template_path=$1
    local target_path=$2

    if [ -f "$target_path" ]; then
        echo "$target_path already exists, leaving it unchanged."
    else
        sed "s/{{YEAR}}/$YEAR/g; s/{{DAY}}/$DAY/g; s/{{DAYPADDED}}/$DAYPADDED/g" "$template_path" > "$target_path"
        echo "Created $target_path from $template_path."
    fi
}

TEMPLATE_DIR="./templates"

create_from_template "$TEMPLATE_DIR/main.go" "$TARGET_DIR/main.go"
create_from_template "$TEMPLATE_DIR/main_test.go" "$TARGET_DIR/main_test.go"
create_from_template "$TEMPLATE_DIR/README.md" "$TARGET_DIR/README.md"
create_from_template "$TEMPLATE_DIR/Makefile" "$TARGET_DIR/Makefile"
create_from_template "$TEMPLATE_DIR/input.txt" "$INPUT_DIR/input.txt"
create_from_template "$TEMPLATE_DIR/sample.txt" "$TARGET_DIR/sample.txt"

# Create or verify symlink to input file in submodule
if [ -L "$TARGET_DIR/input.txt" ] || [ -f "$TARGET_DIR/input.txt" ]; then
    echo "Symlink or file $TARGET_DIR/input.txt already exists."
else
    ln -s "$INPUT_DIR/input.txt" "$TARGET_DIR/input.txt"
    echo "Created symlink $TARGET_DIR/input.txt -> $INPUT_DIR/input.txt"
fi

# Advent of Code 2024 - Day 3

In this readme I detail the puzzle instructions from [adventofcode.com/2024/day/3](https://adventofcode.com/2024/day/3) along with the answers to the puzzle, and any explanation on how I solved it. I typically try to comment the code in such a way that further explanation is not needed, but if things are particularly peculiar then I'll write a little something here.

## Instruction Summary

The North Pole Toboggan Rental Shop's computer memory is corrupted, but you need to determine valid multiplication instructions. The first task is to extract all valid `mul(X,Y)` instructions from the corrupted memory, compute their results, and sum them (e.g., `161` in the example). The second task introduces `do()` and `don't()` instructions, which enable or disable `mul` operations. Only enabled `mul` instructions are processed, requiring a new total based on these rules (e.g., `48` in the example).

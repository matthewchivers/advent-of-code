package main

import (
	"log"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

type lastN struct {
	list []byte
}

func (l *lastN) init(size int) {
	l.list = make([]byte, size)
}

func (l *lastN) add(c byte) {
	l.list = append(l.list[1:], c)
}

func (l *lastN) allUnique() bool {
	chars := make(map[byte]int)
	for _, c := range l.list {
		chars[c]++
		if chars[c] > 1 {
			return false
		}
	}
	return true
}

var (
	data = aoc.ReadFileAsBytes("input.txt")
)

func main() {
	log.Println("Part 1:", partOne())
	log.Println("Part 2:", partTwo())
}

func partOne() int {
	return getFirstUnique(4)
}

func partTwo() int {
	return getFirstUnique(14)
}

func getFirstUnique(size int) int {
	var lastFew lastN
	lastFew.init(size)
	for i, c := range data {
		lastFew.add(c)
		if i > size-1 {
			if lastFew.allUnique() {
				return i + 1
			}
		}
	}
	return -1
}

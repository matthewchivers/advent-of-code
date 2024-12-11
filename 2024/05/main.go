package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/matthewchivers/advent-of-code/util"
)

func main() {
	lines := aoc.ReadFileAsLines("input.txt")
	fmt.Println("Hello, advent of code 2024 - Day 5!")
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
}

// Part One: Sum of middle page numbers from correctly ordered updates
func partOne(input []string) int {
	// Parse rules and updates
	rules, updates := parseInput(input)

	// Process each update
	sumMiddle := 0
	for _, update := range updates {
		if isCorrectOrder(update, rules) {
			middle := getMiddle(update)
			sumMiddle += middle
		}
	}

	return sumMiddle
}

// Part Two: Sum of middle page numbers from reordered incorrectly ordered updates
func partTwo(input []string) int {
	// Parse rules and updates
	rules, updates := parseInput(input)

	// Process each update
	sumMiddle := 0
	for _, update := range updates {
		if !isCorrectOrder(update, rules) {
			sorted, ok := topologicalSort(update, rules)
			if !ok {
				fmt.Println("Error: Unable to sort update:", update)
				continue
			}
			middle := getMiddle(sorted)
			sumMiddle += middle
		}
	}

	return sumMiddle
}

// Helper function to parse the input into rules and updates
func parseInput(input []string) ([][2]int, [][]int) {
	rules := make([][2]int, 0)
	updates := make([][]int, 0)
	section := 0 // 0 for rules, 1 for updates

	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			section = 1
			continue
		}
		if section == 0 {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				continue // skip invalid lines
			}
			x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err1 != nil || err2 != nil {
				continue // skip invalid lines
			}
			rules = append(rules, [2]int{x, y})
		} else {
			pageStrs := strings.Split(line, ",")
			pages := make([]int, 0, len(pageStrs))
			for _, p := range pageStrs {
				num, err := strconv.Atoi(strings.TrimSpace(p))
				if err != nil {
					continue // skip invalid numbers
				}
				pages = append(pages, num)
			}
			if len(pages) > 0 {
				updates = append(updates, pages)
			}
		}
	}

	return rules, updates
}

// Function to check if an update is in the correct order
func isCorrectOrder(update []int, rules [][2]int) bool {
	// Create a map of page to its index for quick lookup
	pageIndex := make(map[int]int)
	for idx, page := range update {
		pageIndex[page] = idx
	}

	// Check each rule
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		idxX, okX := pageIndex[x]
		idxY, okY := pageIndex[y]
		if okX && okY {
			if idxX >= idxY {
				return false
			}
		}
	}
	return true
}

// Function to get the middle page number of an update
func getMiddle(update []int) int {
	n := len(update)
	if n == 0 {
		return 0 // or handle as needed
	}
	mid := n / 2
	if n%2 == 1 {
		return update[mid]
	}
	// If even, return the lower middle as per the example
	return update[mid-1]
}

// Function to perform topological sort on an update based on the rules
func topologicalSort(pages []int, rules [][2]int) ([]int, bool) {
	// Extract relevant rules where both pages are in the update
	relevantRules := make([][2]int, 0)
	pageSet := make(map[int]bool)
	for _, page := range pages {
		pageSet[page] = true
	}
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		if pageSet[x] && pageSet[y] {
			relevantRules = append(relevantRules, rule)
		}
	}

	// Build adjacency list and in-degree map
	adj := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, page := range pages {
		adj[page] = []int{}
		inDegree[page] = 0
	}
	for _, rule := range relevantRules {
		x, y := rule[0], rule[1]
		adj[x] = append(adj[x], y)
		inDegree[y]++
	}

	// Initialize queue with pages having in-degree 0
	// Pages with in-degree 0 have no dependencies and can be processed first.
	queue := make([]int, 0)
	for _, page := range pages {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	// Perform Kahn's algorithm to sort the pages topologically
	// This ensures the sorted order respects all dependencies defined by the rules.
	sorted := make([]int, 0, len(pages))
	for len(queue) > 0 {
		// Remove a page with no dependencies (in-degree 0) from the queue
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		// Reduce the in-degree of its dependent pages
		for _, neighbor := range adj[current] {
			inDegree[neighbor]--
			// If a dependent page now has in-degree 0, add it to the queue
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if topological sort was successful
	// If not all pages are included in the sorted list, it indicates a cycle in the dependencies.
	if len(sorted) != len(pages) {
		return nil, false // Cycle detected or error
	}

	return sorted, true
}

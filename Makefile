# Makefile for Advent of Code project

.PHONY: test test-year test-day benchmark benchmark-year benchmark-day

EXCLUDE_DIR := "advent-of-code/templates"

# Run tests for the entire directory
# The directory context is sensitive: 
#   - Running from root will run all years
#   - Running from a specific day will run only that day
test:
	@echo "Running all tests..."
	@go test $(shell go list ./... | grep -v $(EXCLUDE_DIR))

# Run tests for a specific year (e.g., make test-year YEAR=2020)
# Must be run from the root directory
test-year:
	@echo "Running tests for year $(YEAR)..."
	@go test ./$(YEAR)/...

# Run tests for a specific day (e.g., make test-day YEAR=2020 DAY=day01)
# Must be run from the root directory
test-day:
	@echo "Running tests for year $(YEAR), day $(DAY)..."
	@go test ./$(YEAR)/$(DAY)

# Run benchmarks for the entire directory
# The directory context is sensitive: 
#   - Running from root will run all years
#   - Running from a specific day will run only that day
benchmark:
	@echo "Running all benchmarks..."
	@go test -bench=. ./...

# Run benchmarks for a specific year
# Must be run from the root directory
benchmark-year:
	@echo "Running benchmarks for year $(YEAR)..."
	@go test -bench=. ./$(YEAR)/...

# Run benchmarks for a specific day
# Must be run from the root directory
benchmark-day:
	@echo "Running benchmarks for year $(YEAR), day $(DAY)..."
	@go test -bench=. ./$(YEAR)/$(DAY)

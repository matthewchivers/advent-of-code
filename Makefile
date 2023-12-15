# Makefile for Advent of Code project

.PHONY: test test-year test-day

# Default target: run tests for the entire project
test:
	@echo "Running all tests..."
	@go test ./...

# Run tests for a specific year (e.g., make test-year YEAR=2020)
test-year:
	@echo "Running tests for year $(YEAR)..."
	@go test ./$(YEAR)/...

# Run tests for a specific day (e.g., make test-day YEAR=2020 DAY=day01)
test-day:
	@echo "Running tests for year $(YEAR), day $(DAY)..."
	@go test ./$(YEAR)/$(DAY)

benchmark:
	@echo "Running all benchmarks..."
	@go test -bench=. ./...

benchmark-year:
	@echo "Running benchmarks for year $(YEAR)..."
	@go test -bench=. ./$(YEAR)/...

benchmark-day:
	@echo "Running benchmarks for year $(YEAR), day $(DAY)..."
	@go test -bench=. ./$(YEAR)/$(DAY)
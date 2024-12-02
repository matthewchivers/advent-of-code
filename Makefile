# Makefile for Advent of Code project

.PHONY: test test-year test-day benchmark benchmark-year benchmark-day

EXCLUDE_DIR := "templates"

# Install golangci-lint
setup:
	@echo "Installing golangci-lint..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest || { echo "Failed to install golangci-lint."; exit 1; }
	
# Lint the entire directory
lint:
	@echo "Running golangci-lint on included directories..."
	@golangci-lint run $(shell find . -type f -name '*.go' -not -path "*/templates*" -exec dirname {} \; | sort -u) || { echo "Linting failed. Fix the errors and retry."; exit 1; }

# Run tests for the entire directory
# The directory context is sensitive: 
#   - Running from root will run all years
#   - Running from a specific day will run only that day
test: lint
	@echo "Running all tests..."
	@go test $(shell go list ./... | grep -v $(EXCLUDE_DIR))

# Run tests for a specific year (e.g., make test-year YEAR=2020)
# Must be run from the root directory
test-year: lint
	@echo "Running tests for year $(YEAR)..."
	@go test ./$(YEAR)/...

# Run tests for a specific day (e.g., make test-day YEAR=2020 DAY=day01)
# Must be run from the root directory
test-day: lint
	@echo "Running tests for year $(YEAR), day $(DAY)..."
	@go test ./$(YEAR)/$(DAY)

# Run benchmarks for the entire directory
# The directory context is sensitive: 
#   - Running from root will run all years
#   - Running from a specific day will run only that day
benchmark: lint
	@echo "Running all benchmarks..."
	@go test -bench=. $(shell go list ./... | grep -v $(EXCLUDE_DIR))

benchmark-30: lint
	@echo "Running all benchmarks with 30s benchtime..."
	@go test -bench=. $(shell go list ./... | grep -v $(EXCLUDE_DIR)) -benchtime=30s

# Run benchmarks for a specific year
# Must be run from the root directory
benchmark-year: lint
	@echo "Running benchmarks for year $(YEAR)..."
	@go test -bench=. ./$(YEAR)/...

# Run benchmarks for a specific day
# Must be run from the root directory
benchmark-day: lint
	@echo "Running benchmarks for year $(YEAR), day $(DAY)..."
	@go test -bench=. ./$(YEAR)/$(DAY)

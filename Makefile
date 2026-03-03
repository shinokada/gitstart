.PHONY: build test clean clean-cache clean-all run lint lint-fix coverage install

# Build the gitstart binary
build:
	go build -o gitstart main.go

# Run all tests with verbose output
test:
	go test -v ./...

# Run linter
lint:
	golangci-lint run ./...

# Run linter and auto-fix issues
lint-fix:
	golangci-lint run --fix ./...

# Generate and view coverage report
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Run the CLI tool
run:
	go run main.go

# Clean build artifacts
clean:
	rm -f gitstart coverage.out

# Clean Go caches (test cache and build cache)
clean-cache:
	go clean -testcache -cache

# Clean everything (build artifacts + caches)
clean-all: clean clean-cache

# Install the CLI tool
install:
	go install .

.PHONY: all build test clean clean-cache clean-all run lint lint-fix coverage coverage-ci install ci test-clean test-smoke

# Default target
all: build

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

# Generate and view coverage report (opens browser)
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Generate coverage report (CI-friendly, text output)
coverage-ci:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

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

# Clean, lint, test, and build
ci: clean-all lint test build

# Install the CLI tool
install:
	go install .

# Remove temporary smoke-test directories created under /tmp
test-clean:
	rm -rf /tmp/fake-svelte /tmp/fake-repo /tmp/smoke-node /tmp/smoke-go /tmp/smoke-svelte
	@echo "Cleaned up /tmp smoke-test directories."

# Dry-run smoke tests for all v1.2.0 features (no GitHub required)
test-smoke: install
	@echo "--- Test: --no-license ---"
	gitstart -d /tmp/smoke-node --no-license --dry-run
	@echo "--- Test: --no-readme ---"
	gitstart -d /tmp/smoke-node --no-readme --dry-run
	@echo "--- Test: --post-framework (Node auto-detect) ---"
	mkdir -p /tmp/fake-svelte && touch /tmp/fake-svelte/package.json
	gitstart -d /tmp/fake-svelte --post-framework --dry-run
	@echo "--- Test: branch auto-detection ---"
	mkdir -p /tmp/fake-repo/.git && echo 'ref: refs/heads/develop' > /tmp/fake-repo/.git/HEAD
	touch /tmp/fake-repo/package.json
	gitstart -d /tmp/fake-repo --post-framework --dry-run
	@echo "--- Cleaning up ---"
	$(MAKE) test-clean

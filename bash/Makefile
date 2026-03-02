.PHONY: help test test-unit test-integration test-shellcheck install-deps clean

# Default target
help:
	@echo "Gitstart - Makefile targets"
	@echo ""
	@echo "Testing:"
	@echo "  make test              - Run all tests (shellcheck + unit tests)"
	@echo "  make test-unit         - Run only unit tests"
	@echo "  make test-shellcheck   - Run only shellcheck"
	@echo "  make test-integration  - Run integration tests (requires GitHub)"
	@echo ""
	@echo "Dependencies:"
	@echo "  make install-deps      - Install test dependencies"
	@echo "  make check-deps        - Check if dependencies are installed"
	@echo ""
	@echo "Utilities:"
	@echo "  make clean             - Clean up test artifacts"
	@echo "  make lint              - Run all linters"
	@echo "  make format            - Format shell scripts"
	@echo ""
	@echo "Installation:"
	@echo "  make install           - Install gitstart to /usr/local/bin"
	@echo "  make uninstall         - Uninstall gitstart"

# Run all tests
test: test-shellcheck test-unit
	@echo ""
	@echo "✓ All tests completed!"

# Run shellcheck
test-shellcheck:
	@echo "Running shellcheck..."
	@chmod +x tests/shellcheck.sh
	@./tests/shellcheck.sh

# Run unit tests
test-unit:
	@echo "Running unit tests..."
	@chmod +x tests/*.bats
	@bats tests/gitstart.bats

# Run integration tests
test-integration:
	@echo "Running integration tests..."
	@echo "⚠️  Warning: This will create actual GitHub repositories!"
	@read -p "Continue? [y/N] " -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		bats tests/integration.bats; \
	fi

# Check dependencies
check-deps:
	@echo "Checking dependencies..."
	@command -v shellcheck >/dev/null 2>&1 || echo "❌ shellcheck not found"
	@command -v bats >/dev/null 2>&1 || echo "❌ bats not found"
	@command -v gh >/dev/null 2>&1 || echo "⚠️  gh not found (optional)"
	@command -v jq >/dev/null 2>&1 || echo "⚠️  jq not found (required for gitstart)"
	@echo "Dependency check complete"

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	@if command -v brew >/dev/null 2>&1; then \
		echo "Using Homebrew..."; \
		brew install shellcheck bats-core jq gh; \
	elif command -v apt-get >/dev/null 2>&1; then \
		echo "Using apt-get..."; \
		sudo apt-get update; \
		sudo apt-get install -y shellcheck bats jq gh; \
	else \
		echo "❌ No supported package manager found"; \
		echo "Please install manually:"; \
		echo "  - shellcheck: https://github.com/koalaman/shellcheck"; \
		echo "  - bats: https://github.com/bats-core/bats-core"; \
		echo "  - jq: https://stedolan.github.io/jq/"; \
		echo "  - gh: https://cli.github.com/"; \
		exit 1; \
	fi

# Clean test artifacts
clean:
	@echo "Cleaning test artifacts..."
	@rm -rf tests/*.log
	@rm -rf tests/*.xml
	@rm -rf tests/test-*
	@find . -name "*.bats~" -delete
	@echo "✓ Cleaned"

# Lint all files
lint: test-shellcheck
	@echo "Running additional linters..."
	@if command -v shfmt >/dev/null 2>&1; then \
		echo "Checking formatting with shfmt..."; \
		shfmt -d -i 4 -ci gitstart || echo "⚠️  Formatting issues found"; \
	fi

# Format shell scripts
format:
	@echo "Formatting shell scripts..."
	@if command -v shfmt >/dev/null 2>&1; then \
		shfmt -w -i 4 -ci gitstart; \
		echo "✓ Formatted"; \
	else \
		echo "❌ shfmt not found. Install with: go install mvdan.cc/sh/v3/cmd/shfmt@latest"; \
	fi

# Install gitstart
install:
	@echo "Installing gitstart to /usr/local/bin..."
	@chmod +x gitstart
	@sudo cp gitstart /usr/local/bin/gitstart
	@echo "✓ Installed gitstart to /usr/local/bin/gitstart"
	@echo ""
	@echo "Run 'gitstart -h' to get started"

# Uninstall gitstart
uninstall:
	@echo "Uninstalling gitstart..."
	@chmod +x uninstall.sh
	@./uninstall.sh

# Run full test suite (same as default test runner)
test-all:
	@chmod +x tests/run-tests.sh
	@./tests/run-tests.sh

# Quick test (fast tests only)
test-quick: test-shellcheck
	@echo "Running quick tests..."
	@bats tests/gitstart.bats --filter "version\|help\|dry-run"

# Continuous testing (watch mode)
watch:
	@echo "Watching for changes..."
	@echo "Note: Install 'entr' for file watching"
	@if command -v entr >/dev/null 2>&1; then \
		find . -name "*.sh" -o -name "*.bats" -o -name "gitstart" | entr -c make test-quick; \
	else \
		echo "Install entr: brew install entr (macOS) or apt install entr (Linux)"; \
	fi

# Code coverage (approximate)
coverage:
	@echo "Test coverage analysis..."
	@echo "Total functions in gitstart:"
	@grep -c "^[a-zA-Z_][a-zA-Z0-9_]*() {" gitstart || echo "0"
	@echo "Total test cases:"
	@grep -c "^@test" tests/gitstart.bats || echo "0"
	@echo ""
	@echo "Note: This is an approximate count. For detailed coverage, use coverage tools."

# Pre-commit checks
pre-commit: test-shellcheck test-quick
	@echo "✓ Pre-commit checks passed"

# Create release
release:
	@echo "Creating release..."
	@echo "Current version: $$(./gitstart -v)"
	@echo ""
	@echo "Steps for release:"
	@echo "1. Update version in gitstart script"
	@echo "2. Update CHANGELOG.md"
	@echo "3. Run: make test"
	@echo "4. Commit changes"
	@echo "5. Tag release: git tag -a vX.Y.Z -m 'Release X.Y.Z'"
	@echo "6. Push: git push && git push --tags"

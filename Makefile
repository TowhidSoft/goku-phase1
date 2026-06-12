BINARY    := goku
BUILD_DIR := ./bin
PKG       := github.com/yourusername/goku

.PHONY: all build test lint clean run help

## all: build the binary
all: build

## build: compile the binary into ./bin/goku
build:
	@echo "→ Building $(BINARY)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY) ./main.go
	@echo "✓ Binary ready at $(BUILD_DIR)/$(BINARY)"

## test: run all unit tests
test:
	@echo "→ Running tests..."
	go test ./... -v -count=1

## lint: run go vet + staticcheck
lint:
	@echo "→ Running linters..."
	go vet ./...

## clean: remove build artifacts
clean:
	@echo "→ Cleaning..."
	@rm -rf $(BUILD_DIR)

## run: build and show help
run: build
	$(BUILD_DIR)/$(BINARY) --help

## help: list available make targets
help:
	@grep -E '^## ' Makefile | sed 's/## //'

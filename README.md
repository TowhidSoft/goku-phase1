# goku phase 1 ⚡

A blazing-fast CLI tool for converting configuration files between **JSON** and **YAML** formats, built with Go and [Cobra](https://github.com/spf13/cobra).

---

## Project Structure

```
goku/
├── main.go                        # Entry point
├── go.mod / go.sum
├── Makefile                       # Developer workflow
├── examples/
│   ├── config.json
│   └── config.yaml
└── cmd/                           # Cobra CLI layer
│   ├── root.go                    # Root command + Execute()
│   ├── convert.go                 # `goku convert` subcommand
│   └── version.go                 # `goku version` subcommand
└── internal/                      # Business logic (not importable externally)
    ├── config/
    │   └── loader.go              # Read file → GenericConfig
    ├── converter/
    │   ├── converter.go           # GenericConfig → target format string
    │   └── converter_test.go
    └── validator/
        ├── validator.go           # Flag validation, format detection, edge cases
        └── validator_test.go
```

---

## Installation

```bash
git clone https://github.com/towhidSoft/goku
cd goku
make build
# Binary is at ./bin/goku
```

---

## Usage

```bash
# JSON → YAML
goku convert -i config.json -o yaml

# YAML → JSON
goku convert -i config.yaml -o json

# Check version
goku version
```

### Edge Case — Same Format Error

```bash
goku convert -i config.json -o json
# Error: requested output must be in a different format than the input file
```

---

## Development

```bash
make test    # Run all unit tests
make lint    # Run go vet
make build   # Compile binary
make clean   # Remove build artifacts
```

---

## Phase 1 Checklist

- [x] Cobra CLI integration (`goku convert`, `goku version`)
- [x] Read configuration file (JSON / YAML)
- [x] Map into Go struct (`GenericConfig`)
- [x] Convert and print output
- [x] Edge case: same input/output format returns a clear error
- [x] Unit tests for validator and converter
- [x] Production-grade package layout (`cmd/`, `internal/`)

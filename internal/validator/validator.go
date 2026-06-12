package validator

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// SupportedFormats lists all formats goku understands.
var SupportedFormats = map[string]bool{
	"json": true,
	"yaml": true,
	"yml":  true,
}

// ErrSameFormat is returned when input and output formats are identical.
var ErrSameFormat = errors.New("requested output must be in a different format than the input file")

// ValidateFlags performs basic sanity checks on the CLI flags.
func ValidateFlags(inputFile, outputFormat string) error {
	if strings.TrimSpace(inputFile) == "" {
		return errors.New("input file path cannot be empty")
	}

	normalized := normalizeFormat(outputFormat)
	if !SupportedFormats[normalized] {
		return fmt.Errorf("unsupported output format %q — choose json or yaml", outputFormat)
	}

	return nil
}

// DetectFormat infers the file format from the file extension.
func DetectFormat(filePath string) (string, error) {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filePath), "."))

	switch ext {
	case "json":
		return "json", nil
	case "yaml", "yml":
		return "yaml", nil
	default:
		return "", fmt.Errorf("unsupported file extension %q — expected .json, .yaml, or .yml", ext)
	}
}

// EnsureFormatsDiffer enforces the rule that input and output formats must differ.
func EnsureFormatsDiffer(inputFormat, outputFormat string) error {
	if normalizeFormat(inputFormat) == normalizeFormat(outputFormat) {
		return ErrSameFormat
	}
	return nil
}

// normalizeFormat maps "yml" → "yaml" so comparisons are canonical.
func normalizeFormat(f string) string {
	f = strings.ToLower(strings.TrimSpace(f))
	if f == "yml" {
		return "yaml"
	}
	return f
}

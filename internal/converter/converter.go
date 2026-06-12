package converter

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/towhidSoft/goku/internal/config"
)

// Convert marshals a GenericConfig into the requested output format string.
func Convert(cfg config.GenericConfig, outputFormat string) (string, error) {
	switch outputFormat {
	case "json":
		return toJSON(cfg)
	case "yaml":
		return toYAML(cfg)
	default:
		return "", fmt.Errorf("unsupported output format: %q", outputFormat)
	}
}

func toJSON(cfg config.GenericConfig) (string, error) {
	// Use indented JSON for human-readable output.
	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(out) + "\n", nil
}

func toYAML(cfg config.GenericConfig) (string, error) {
	out, err := yaml.Marshal(cfg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal to YAML: %w", err)
	}
	return string(out), nil
}

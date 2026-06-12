package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// GenericConfig is a format-agnostic representation of the config file.
// Using map[string]any lets us handle arbitrary nesting without
// defining a fixed schema — the converter re-marshals it as needed.
type GenericConfig map[string]any

// Load reads a file from disk and parses it according to the detected format.
func Load(filePath, format string) (GenericConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %q: %w", filePath, err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file %q is empty", filePath)
	}

	switch format {
	case "json":
		return parseJSON(data)
	case "yaml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("unsupported format: %q", format)
	}
}

func parseJSON(data []byte) (GenericConfig, error) {
	var cfg GenericConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return cfg, nil
}

func parseYAML(data []byte) (GenericConfig, error) {
	var cfg GenericConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}
	return cfg, nil
}

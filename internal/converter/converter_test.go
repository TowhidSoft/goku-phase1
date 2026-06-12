package converter_test

import (
	"strings"
	"testing"

	"github.com/towhidSoft/goku/internal/config"
	"github.com/towhidSoft/goku/internal/converter"
)

var sampleCfg = config.GenericConfig{
	"name":    "goku",
	"version": "1.0.0",
	"server": map[string]any{
		"host": "localhost",
		"port": 8080,
	},
}

func TestConvertToYAML(t *testing.T) {
	out, err := converter.Convert(sampleCfg, "yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "name:") {
		t.Errorf("expected YAML output to contain 'name:', got:\n%s", out)
	}
	if !strings.Contains(out, "goku") {
		t.Errorf("expected YAML output to contain 'goku', got:\n%s", out)
	}
}

func TestConvertToJSON(t *testing.T) {
	out, err := converter.Convert(sampleCfg, "json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `"name"`) {
		t.Errorf("expected JSON output to contain '\"name\"', got:\n%s", out)
	}
	if !strings.Contains(out, "goku") {
		t.Errorf("expected JSON output to contain 'goku', got:\n%s", out)
	}
}

func TestConvertUnsupportedFormat(t *testing.T) {
	_, err := converter.Convert(sampleCfg, "toml")
	if err == nil {
		t.Errorf("expected error for unsupported format, got nil")
	}
}

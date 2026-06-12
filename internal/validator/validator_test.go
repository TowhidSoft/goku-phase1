package validator_test

import (
	"errors"
	"testing"

	"github.com/towhidSoft/goku/internal/validator"
)

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		name       string
		filePath   string
		wantFormat string
		wantErr    bool
	}{
		{"json file", "config.json", "json", false},
		{"yaml file", "config.yaml", "yaml", false},
		{"yml file", "config.yml", "yaml", false},
		{"uppercase extension", "config.JSON", "json", false},
		{"unsupported extension", "config.toml", "", true},
		{"no extension", "configfile", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := validator.DetectFormat(tc.filePath)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tc.wantFormat {
				t.Errorf("got %q, want %q", got, tc.wantFormat)
			}
		})
	}
}

func TestEnsureFormatsDiffer(t *testing.T) {
	tests := []struct {
		name         string
		inputFormat  string
		outputFormat string
		wantErr      bool
	}{
		{"json to yaml — ok", "json", "yaml", false},
		{"yaml to json — ok", "yaml", "json", false},
		{"json to json — error", "json", "json", true},
		{"yaml to yaml — error", "yaml", "yaml", true},
		{"yaml to yml — error (same canonical)", "yaml", "yml", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.EnsureFormatsDiffer(tc.inputFormat, tc.outputFormat)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				if !errors.Is(err, validator.ErrSameFormat) {
					t.Errorf("expected ErrSameFormat, got %v", err)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestValidateFlags(t *testing.T) {
	tests := []struct {
		name         string
		inputFile    string
		outputFormat string
		wantErr      bool
	}{
		{"valid json output", "config.yaml", "json", false},
		{"valid yaml output", "config.json", "yaml", false},
		{"empty input", "", "yaml", true},
		{"unsupported output format", "config.json", "toml", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateFlags(tc.inputFile, tc.outputFormat)
			if tc.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

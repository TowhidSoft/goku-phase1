package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/towhidSoft/goku/internal/config"
	"github.com/towhidSoft/goku/internal/converter"
	"github.com/towhidSoft/goku/internal/validator"
)

var (
	inputFile    string
	outputFormat string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a config file between JSON and YAML",
	Long: `Convert reads a JSON or YAML configuration file and outputs
the converted content in the requested format.

The input and output formats must differ.

Examples:
  goku convert -i config.json -o yaml
  goku convert -i config.yaml -o json`,

	RunE: runConvert,
}

func init() {
	convertCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the input config file (JSON or YAML) [required]")
	convertCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Desired output format: json or yaml [required]")

	_ = convertCmd.MarkFlagRequired("input")
	_ = convertCmd.MarkFlagRequired("output")
}

func runConvert(cmd *cobra.Command, args []string) error {
	// 1. Validate flags
	if err := validator.ValidateFlags(inputFile, outputFormat); err != nil {
		return fmt.Errorf("%w", err)
	}

	// 2. Detect input format from file extension
	inputFormat, err := validator.DetectFormat(inputFile)
	if err != nil {
		return fmt.Errorf("cannot determine input format: %w", err)
	}

	// 3. Edge-case: same input and output format
	if err := validator.EnsureFormatsDiffer(inputFormat, outputFormat); err != nil {
		return fmt.Errorf("%w", err)
	}

	// 4. Load & parse the config file into a generic map
	cfg, err := config.Load(inputFile, inputFormat)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	// 5. Convert to the target format
	result, err := converter.Convert(cfg, outputFormat)
	if err != nil {
		return fmt.Errorf("conversion failed: %w", err)
	}

	// 6. Print output
	fmt.Fprint(os.Stdout, result)
	return nil
}

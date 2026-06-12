package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "Goku — a blazing-fast config file converter (JSON ↔ YAML)",
	Long: `
  ______     ______     __  __     __  __    
 /\  ___\   /\  __ \   /\ \/ /    /\ \/\ \   
 \ \ \__ \  \ \ \/\ \  \ \  _"-.  \ \ \_\ \  
  \ \_____\  \ \_____\  \ \_\ \_\  \ \_____\ 
   \/_____/   \/_____/   \/_/\/_/   \/_____/ 

Goku converts configuration files between JSON and YAML formats.

Usage:
  goku convert -i <input_file> -o <output_format>

Examples:
  goku convert -i config.json -o yaml
  goku convert -i config.yaml -o json
`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute is the entry point called by main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(versionCmd)
}

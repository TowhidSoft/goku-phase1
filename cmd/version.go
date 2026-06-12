package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version   = "1.0.0"
	BuildDate = "2026-06-13"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of goku",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("goku v%s (built %s)\n", Version, BuildDate)
	},
}

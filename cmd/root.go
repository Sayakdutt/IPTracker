package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "iptracker",
		Short: "IPTRACKER CLI APPLICATION",
		Long:  `IPTRACKER CLI APPLICATION`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

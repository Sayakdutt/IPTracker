/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the Version",
	Long:  `Get the Version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("v1.0.1")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

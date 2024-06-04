/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/timzolleis/hfu-ain-pas/cli/utils"

	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy files to your target",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := args[1]
		config := utils.ReadConfig()
		utils.CopyFiles(cmd, config, source, destination)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}

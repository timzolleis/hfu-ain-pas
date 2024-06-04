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
		target := cmd.Flag("target").Value.String()
		ip := cmd.Flag("ip").Value.String()
		utils.CopyFiles(cmd, &utils.SSHConfig{
			Host:     target,
			User:     cmd.Flag("user").Value.String(),
			Password: cmd.Flag("password").Value.String(),
			Ip:       ip,
		}, source, destination)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("ip", "i", "", "The ip address to connect to")
	copyCmd.Flags().StringP("target", "t", "target068", "The target to copy to")
	copyCmd.Flags().String("user", "pi", "The user to connect as")
	copyCmd.Flags().String("password", "raspberry", "The password to use")
}

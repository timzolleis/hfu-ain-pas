/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timzolleis/hfu-ain-pas/cli/utils"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the raspberry pi",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		utils.ConnectSSH(cmd, config)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

}

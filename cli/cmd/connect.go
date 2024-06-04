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
		username, _ := cmd.Flags().GetString("user")
		target, _ := cmd.Flags().GetString("target")
		password, _ := cmd.Flags().GetString("password")
		ip := cmd.Flag("ip").Value.String()
		utils.ConnectSSH(cmd, &utils.SSHConfig{
			Host:     target,
			User:     username,
			Password: password,
			Ip:       ip,
		})
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringP("target", "t", "target068", "The target to connect to")
	connectCmd.Flags().StringP("ip", "i", "", "The ip address to connect to")
	connectCmd.Flags().StringP("user", "u", "pi", "The user to connect as")
	connectCmd.Flags().StringP("password", "p", "raspberry", "The password to use")

}

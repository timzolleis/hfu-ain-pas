/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timzolleis/hfu-ain-pas/cli/utils"
)

// connectHostCmd represents the connectHost command
var connectHostCmd = &cobra.Command{
	Use:   "connect-host",
	Short: "Connect to the Host via SSH",
	Run: func(cmd *cobra.Command, args []string) {
		target := cmd.Flag("target").Value.String()
		user := cmd.Flag("user").Value.String()
		password := cmd.Flag("password").Value.String()
		utils.ConnectSSH(cmd, &utils.SSHConfig{
			Host:     target,
			User:     user,
			Password: password,
		})
	},
}

func init() {
	rootCmd.AddCommand(connectHostCmd)
	connectHostCmd.Flags().StringP("target", "t", "host068", "The target to connect to")
	connectHostCmd.Flags().StringP("user", "u", "tim", "The user to connect as")
	connectHostCmd.Flags().StringP("password", "p", "tim", "The password to use")

}

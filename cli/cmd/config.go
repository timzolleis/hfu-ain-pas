/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timzolleis/hfu-ain-pas/cli/utils"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Change the config",
	Run: func(cmd *cobra.Command, args []string) {
		currentConf := utils.ReadConfig()
		newConf := utils.QueryConfig()
		if newConf.Target != currentConf.Target || newConf.User != currentConf.User {
			utils.SetupSsh(newConf)
		}
		utils.WriteConfig(newConf)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

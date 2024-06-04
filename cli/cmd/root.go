/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/timzolleis/hfu-ain-pas/cli/utils"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("target", "t", "", "The target to connect to")
	rootCmd.PersistentFlags().StringP("user", "u", "", "The user to connect as")
	rootCmd.PersistentFlags().StringP("password", "p", "", "The password to use")
	setup()
}

func setup() {
	var config *utils.Config
	if !utils.HasConfig() {
		config = utils.QueryConfig()
	} else {
		config = utils.ReadConfig()
	}
	utils.UpdateConfig(rootCmd, config)
	if !config.KeyUploaded {
		utils.SetupSsh(config)
	}

}

package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

func CopyFiles(cmd *cobra.Command, config *SSHConfig, source, destination string) {
	if config.Ip != "" {
		config.Host = config.Ip
	} else {
		config.Host += ".local"
	}
	scpString := fmt.Sprintf("%s@%s:%s", config.User, config.Host, destination)
	command := exec.Command("scp", "-r", source, scpString)
	command.Stderr = cmd.OutOrStdout()
	command.Stdout = cmd.OutOrStderr()
	err := command.Run()
	if err != nil {
		log.Fatalf("Error running scp command: %s", err)
	}
}

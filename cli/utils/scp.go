package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

func CopyFiles(cmd *cobra.Command, config *Config, source, destination string) {
	scpString := fmt.Sprintf("%s@%s:%s", config.User, config.Target, destination)
	command := exec.Command("scp", "-r", source, scpString)
	command.Stderr = cmd.OutOrStdout()
	command.Stdout = cmd.OutOrStderr()
	err := command.Run()
	if err != nil {
		log.Fatalf("Error running scp command: %s", err)
	}
}

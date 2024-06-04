package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

type SSHConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}

func ConnectSSH(cmd *cobra.Command, config *SSHConfig) {
	if config.Ip != "" {
		log.Println("Connecting to ip")
		config.Host = config.Ip
	} else {
		log.Println("Connecting to local")
		config.Host += ".local"
	}
	connectString := fmt.Sprintf("%s@%s", config.User, config.Host)
	command := exec.Command("ssh", "-o", "PubkeyAuthentication=no", connectString)
	command.Stdout = cmd.OutOrStdout()
	command.Stderr = cmd.OutOrStderr()
	command.Stdin = cmd.InOrStdin()
	err := command.Run()
	if err != nil {
		log.Fatalf("Error running ssh command: %s", err)
	}
}

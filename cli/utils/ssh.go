package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type SSHConfig struct {
	Host    string `json:"host"`
	User    string `json:"user"`
	KeyPath string `json:"keyPath"`
}

func ConnectSSH(cmd *cobra.Command, config *Config) {
	connectString := fmt.Sprintf("%s@%s", config.User, config.Target)
	command := exec.Command("ssh", "-i", config.KeyFile, "-o", "IdentitiesOnly=yes", connectString)
	command.Stdout = cmd.OutOrStdout()
	command.Stderr = cmd.OutOrStderr()
	command.Stdin = cmd.InOrStdin()
	err := command.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok && status.ExitStatus() == 130 {
				return
			}
		}
		log.Fatalf("Error running ssh command: %s", err)
	}
}

func UploadKey(config *Config, publicPath string) {
	log.Println("Uploading key to target", publicPath)
	publicKey, err := os.ReadFile(publicPath)
	if err != nil {
		log.Fatalf("Error reading public key file: %s", err)
	}
	command := exec.Command("ssh", "-o", "PubkeyAuthentication=no", fmt.Sprintf("%s@%s", config.User, config.Target), fmt.Sprintf("echo '%s' >> ~/.ssh/authorized_keys", string(publicKey)))
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		log.Fatalf("Error running ssh command: %s", err)
	}
}

func SetupSsh(config *Config) {
	log.Println("Uploading key...")
	//Generate key
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		log.Fatalf("Error generating key pair: %s", err)
	}
	//Write the key to the file system
	privatePath, publicPath := WriteKeyPair(privateKey, publicKey)
	config.KeyFile = privatePath
	//Upload the key to the target
	UploadKey(config, publicPath)
	config.KeyUploaded = true
	WriteConfig(config)
}

package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	User        string `json:"user"`
	Target      string `json:"target"`
	KeyFile     string `json:"keyFile"`
	KeyUploaded bool   `json:"keyUploaded"`
}

func getConfigPath() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return filepath.Join(homedir, ".piconfig/config.json")
}

func makeDirectory() {
	dir := filepath.Dir(getConfigPath())
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
		log.Println("Created directory: ", dir)
	}
}

func HasConfig() bool {
	_, err := os.Stat(getConfigPath())
	return err == nil
}

func ReadConfig() *Config {
	content, err := os.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	config := Config{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return &config
}

func WriteConfig(config *Config) {
	makeDirectory()
	content, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	err = os.WriteFile(getConfigPath(), content, 0644)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func UpdateConfig(cmd *cobra.Command, config *Config) {
	username, _ := cmd.PersistentFlags().GetString("user")
	target, _ := cmd.PersistentFlags().GetString("target")
	if username != "" {
		config.User = username
	}
	if target != "" {
		config.Target = target
	}
	WriteConfig(config)
}

func WriteKeyPair(privateKey *rsa.PrivateKey, publicKey []byte) (string, string) {
	privatePath := filepath.Join(filepath.Dir(getConfigPath()), "key")
	publicPath := filepath.Join(filepath.Dir(getConfigPath()), "key.pub")
	privatePEM := encodePrivateKeyToPEM(privateKey)
	err := os.WriteFile(privatePath, privatePEM, 0600)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	err = os.WriteFile(publicPath, publicKey, 0644)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	err = os.Chmod(privatePath, 0600)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return privatePath, publicPath
}

func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

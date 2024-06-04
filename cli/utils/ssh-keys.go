package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"golang.org/x/crypto/ssh"
	"log"
)

func GenerateKeyPair() (privateKey *rsa.PrivateKey, publicKey []byte, err error) {
	bitSize := 4096
	privateKey, err = generatePrivateKey(bitSize)
	if err != nil {
		return nil, nil, err
	}
	publicKey, err = generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, publicKey, nil
}

func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	log.Println("INFO: Generating Private Key")
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}
	log.Println("SUCCESS: Private Key generated")
	return privateKey, nil
}

func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	log.Println("INFO: Generating Public Key")
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)
	log.Println("SUCCESS: Public Key generated")
	return pubKeyBytes, nil
}

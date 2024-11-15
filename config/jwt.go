package config

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

func initJwt() {
	privateKeyBytes, err := os.ReadFile("key_jwt/private_key.pem")
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatalf("Failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse RSA private key: %v", err)
	}

	jwt = jwtauth.New("RS256", privateKey, nil)
}

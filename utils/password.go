package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/argon2"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) []byte {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return hash
}

func HashPasswordToBase64(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		log.Println("Error generating salt:", err)
		return "", err
	}

	hashedPassword := hashPassword(password, salt)
	hashedPasswordStr := base64.StdEncoding.EncodeToString(hashedPassword)

	return hashedPasswordStr, nil
}

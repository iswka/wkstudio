package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// EncryptPassword encrypts a password using AES-256-GCM
func EncryptPassword(plaintext, key string) (string, error) {
	// Convert key to 32 bytes (AES-256 requires 32-byte key)
	keyBytes := make([]byte, 32)
	copy(keyBytes, []byte(key))
	if len(key) < 32 {
		// If key is too short, pad by repeating
		for i := len(key); i < 32; i++ {
			keyBytes[i] = keyBytes[i%len(key)]
		}
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptPassword decrypts a password
func DecryptPassword(ciphertext, key string) (string, error) {
	// Convert key to 32 bytes
	keyBytes := make([]byte, 32)
	copy(keyBytes, []byte(key))
	if len(key) < 32 {
		for i := len(key); i < 32; i++ {
			keyBytes[i] = keyBytes[i%len(key)]
		}
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(data) < gcm.NonceSize() {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertextBytes := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

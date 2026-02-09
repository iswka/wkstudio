package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// EncryptPassword 加密密码（使用AES-256-GCM）
func EncryptPassword(plaintext, key string) (string, error) {
	// 将key转换为32字节（AES-256需要32字节密钥）
	keyBytes := make([]byte, 32)
	copy(keyBytes, []byte(key))
	if len(key) < 32 {
		// 如果key太短，重复填充
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

// DecryptPassword 解密密码
func DecryptPassword(ciphertext, key string) (string, error) {
	// 将key转换为32字节
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

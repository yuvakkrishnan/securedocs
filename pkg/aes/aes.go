package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// FixKey adjusts the key to the correct length (16, 24, or 32 bytes).
func fixKey(key string, length int) string {
	if len(key) > length {
		return key[:length]
	}
	for len(key) < length {
		key += "0"
	}
	return key
}

// Encrypt encrypts plaintext using AES.
func Encrypt(plaintext, key string) (string, error) {
	// Adjust the key to 32 bytes for AES-256
	key = fixKey(key, 32)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	plaintextBytes := []byte(plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextBytes)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES.
func Decrypt(ciphertext, key string) (string, error) {
	// Adjust the key to 32 bytes for AES-256
	key = fixKey(key, 32)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	if len(ciphertextBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)
	return string(ciphertextBytes), nil
}

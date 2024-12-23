package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
	"securedocs/internal/utils"
)

// EncryptFile encrypts a binary file using AES and saves the output.
func EncryptFile(inputPath, outputPath, key string) error {
	// Fix the key length for AES-256
	key = utils.FixKey(key, 32)

	// Read the input file
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Create AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// Prepare ciphertext
	ciphertext := make([]byte, aes.BlockSize+len(inputData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], inputData)

	// Write encrypted data to output file
	return os.WriteFile(outputPath, ciphertext, 0644)
}

// DecryptFile decrypts an AES-encrypted binary file and saves the output.
func DecryptFile(inputPath, outputPath, key string) error {
	// Fix the key length for AES-256
	key = utils.FixKey(key, 32)

	// Read the encrypted file
	ciphertext, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Create AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	if len(ciphertext) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	// Write decrypted data to output file
	return os.WriteFile(outputPath, ciphertext, 0644)
}

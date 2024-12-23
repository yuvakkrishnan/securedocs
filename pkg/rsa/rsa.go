package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

// Encrypt encrypts plaintext using RSA public key.
func Encrypt(plaintext string, publicKey *rsa.PublicKey) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(plaintext), nil)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using RSA private key.
func Decrypt(ciphertext string, privateKey *rsa.PrivateKey) (string, error) {
	ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

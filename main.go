package main

import (
	"fmt"
	"encryption-decryption-system/pkg/aes"
	"encryption-decryption-system/pkg/rsa"
	"encryption-decryption-system/pkg/keys"
)

func main() {
	// AES Example
	key := "thisis32bitlongpassphrase!!!" // 32 bytes for AES-256
	plaintext := "Sensitive Document Content"

	encryptedAES, err := aes.Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("Error encrypting AES:", err)
		return
	}
	fmt.Println("AES Encrypted:", encryptedAES)

	decryptedAES, err := aes.Decrypt(encryptedAES, key)
	if err != nil {
		fmt.Println("Error decrypting AES:", err)
		return
	}
	fmt.Println("AES Decrypted:", decryptedAES)

	// RSA Example
	privateKey, publicKey, err := keys.GenerateRSAKeys()
	if err != nil {
		fmt.Println("Error generating RSA keys:", err)
		return
	}

	encryptedRSA, err := rsa.Encrypt(plaintext, publicKey)
	if err != nil {
		fmt.Println("Error encrypting RSA:", err)
		return
	}
	fmt.Println("RSA Encrypted:", encryptedRSA)

	decryptedRSA, err := rsa.Decrypt(encryptedRSA, privateKey)
	if err != nil {
		fmt.Println("Error decrypting RSA:", err)
		return
	}
	fmt.Println("RSA Decrypted:", decryptedRSA)
}

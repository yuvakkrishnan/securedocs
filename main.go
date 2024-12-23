package main

import (
	"fmt"
	"securedocs/pkg/aes"
)

func main() {
	// Define file paths
	inputFilePath := "testdata/sample.pdf"         // Input file
	encryptedFilePath := "testdata/sample.enc"    // Encrypted file
	decryptedFilePath := "testdata/sample_dec.pdf" // Decrypted file

	// AES Key (32 bytes for AES-256)
	key := "thisis32bitlongpassphrase32!"

	// Encrypt the file
	err := aes.EncryptFile(inputFilePath, encryptedFilePath, key)
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}
	fmt.Println("File encrypted successfully:", encryptedFilePath)

	// Decrypt the file
	err = aes.DecryptFile(encryptedFilePath, decryptedFilePath, key)
	if err != nil {
		fmt.Println("Error decrypting file:", err)
		return
	}
	fmt.Println("File decrypted successfully:", decryptedFilePath)
}

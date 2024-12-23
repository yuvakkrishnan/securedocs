package keys

import (
	"crypto/rand"
	"crypto/rsa"
)

// GenerateRSAKeys generates RSA private and public keys.
func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

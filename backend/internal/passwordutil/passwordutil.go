package passwordutil

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPasswordWithSalt(password string, secret string) (string, error) {
	passwordWithSalt := append([]byte(password), []byte(secret)...)

	hash := sha256.Sum256(passwordWithSalt)

	hashString := hex.EncodeToString(hash[:])

	return hashString, nil
}

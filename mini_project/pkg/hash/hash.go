package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// ComparePassword compares a hashed password with a plaintext password
func ComparePassword(hashedPassword, password string) bool {
	return hashedPassword == HashPassword(password)
}

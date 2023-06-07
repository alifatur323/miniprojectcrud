package middleware

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func VerifyPassword(password, hashedPassword string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString == hashedPassword
}

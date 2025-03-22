package jwtgo

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey

// GenerateToken generates a new JWT token
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error in generating token")
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates a JWT token
func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("Error in validating token")
		return false
	}
	if token.Valid {
		return true
	}
	return false
}

// Init initializes the JWT key
func Init(key string) error {
	if key == "" {
		return errors.New("JWT key is empty")
	}
	jwtKey = []byte(key)
	return nil
}

// GetJWTKey returns the JWT key
func GetJWTKey() []byte {
	return jwtKey
}

// SetJWTKey sets the JWT key
func SetJWTKey(key string) {

	jwtKey = []byte(key)
}

package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// In production, load the secret from environment variables
	secret := []byte("your-256-bit-secret")

	// Create claims with standard and custom fields
	claims := jwt.MapClaims{
		"sub":   "1234567890",
		"name":  "John Doe",
		"admin": true,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	fmt.Println("Signed JWT:")
	fmt.Println(tokenString)
}
package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// MyCustomClaims embeds RegisteredClaims and adds custom fields.
type MyCustomClaims struct {
	jwt.RegisteredClaims
	Role  string `json:"role"`
	Scope string `json:"scope"`
}

func main() {
	secret := []byte("custom-claims-secret")

	// --- Create a token with custom claims ---
	claims := MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "my-app",
			Subject:   "user-42",
			Audience:  jwt.ClaimStrings{"web", "mobile"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Role:  "admin",
		Scope: "read:users write:users",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(secret)
	fmt.Println("Token:", signed)

	// --- Parse and validate the token ---
	parsedToken, err := jwt.ParseWithClaims(signed, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}

	if customClaims, ok := parsedToken.Claims.(*MyCustomClaims); ok && parsedToken.Valid {
		fmt.Printf("User %s has role %s with scope %s\n",
			customClaims.Subject, customClaims.Role, customClaims.Scope)
	}
}
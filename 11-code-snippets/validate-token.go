package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateToken parses a token string and performs validation on standard claims.
// keyFunc should return the correct key based on the token's method.
func ValidateToken(tokenString string, keyFunc jwt.Keyfunc, expectedIssuer string, expectedAudience string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, keyFunc,
		jwt.WithLeeway(1*time.Minute), // allow clock skew
		jwt.WithValidMethods([]string{"HS256", "RS256"}),
		jwt.WithIssuer(expectedIssuer),
		jwt.WithAudience(expectedAudience),
	)
	if err != nil {
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected claims type")
	}

	return claims, nil
}
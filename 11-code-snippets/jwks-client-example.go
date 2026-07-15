package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

func main() {
	// Example JWKS endpoint (Auth0, Okta, etc.)
	jwksURL := "https://your-domain.com/.well-known/jwks.json"
	tokenString := "your.jwt.token.here"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch JWKS
	set, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		log.Fatalf("failed to fetch JWKS: %v", err)
	}

	// Parse the JWT without validation first to get the kid
	unverifiedToken, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Fatalf("failed to parse token: %v", err)
	}

	kid, ok := unverifiedToken.Header["kid"].(string)
	if !ok {
		log.Fatal("token header missing kid")
	}

	// Look up the key from the JWKS
	key, found := set.LookupKeyID(kid)
	if !found {
		log.Fatalf("key with kid %s not found in JWKS", kid)
	}

	// Convert JWK to a raw public key
	var pubKey interface{}
	if err := key.Raw(&pubKey); err != nil {
		log.Fatalf("failed to get raw key: %v", err)
	}

	// Verify the token with the obtained key
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure algorithm matches
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return pubKey, nil
	})
	if err != nil || !token.Valid {
		log.Fatalf("token verification failed: %v", err)
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Printf("Token verified. Subject: %v\n", claims["sub"])
}
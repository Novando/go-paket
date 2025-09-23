package jwt

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
)

// ParseToken parses and validates a JWT token and maps it to the provided target struct.
// The target parameter should be a pointer to a struct that embeds IDTokenBase.
func ParseToken(tokenString, secret string, target interface{}) error {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	// Verify token
	if !token.Valid {
		return ErrTokenInvalid
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("failed to parse claims")
	}

	// Convert claims to JSON and then to the target struct
	return mapClaimsToStruct(claims, target)
}

// mapClaimsToStruct converts jwt.MapClaims to the target struct
func mapClaimsToStruct(claims jwt.MapClaims, target interface{}) error {
	// Convert claims to JSON bytes
	jsonData, err := json.Marshal(claims)
	if err != nil {
		return fmt.Errorf("failed to marshal claims: %w", err)
	}

	// Unmarshal JSON into target struct
	if err := json.Unmarshal(jsonData, target); err != nil {
		return fmt.Errorf("failed to unmarshal claims: %w", err)
	}

	// Handle time values which might be float64 in the claims
	if tokenBase, ok := target.(interface {
		SetExpirationTime(t time.Time)
		SetNotBefore(t time.Time)
		SetIssuedAt(t time.Time)
	}); ok {
		if exp, ok := claims["exp"].(float64); ok {
			tokenBase.SetExpirationTime(time.Unix(int64(exp), 0))
		}
		if nbf, ok := claims["nbf"].(float64); ok {
			tokenBase.SetNotBefore(time.Unix(int64(nbf), 0))
		}
		if iat, ok := claims["iat"].(float64); ok {
			tokenBase.SetIssuedAt(time.Unix(int64(iat), 0))
		}
	}

	return nil
}

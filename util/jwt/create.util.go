package jwt

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Create to issue JWT
func Create(
	appUrl string,
	clientUrlOrIp string,
	subject string,
	claims any,
	secret string,
	opt Option,
) (string, error) {
	var (
		signingMethod = jwt.GetSigningMethod("HS256")
		duration      = 2 * time.Hour
	)

	// assign option if exist
	if opt.Method != nil {
		signingMethod = opt.Method
	}
	if opt.Duration != 0 {
		duration = opt.Duration
	}

	// Create a new token
	token := jwt.New(signingMethod)

	// Convert claims to map
	claimsMap, err := toMap(claims)
	if err != nil {
		return "", fmt.Errorf("failed to convert claims: %w", err)
	}

	// Set expiration
	iat := time.Now()
	claimsMap["exp"] = iat.Add(duration).Unix()
	claimsMap["iat"] = iat.Unix()
	claimsMap["nbf"] = iat.Unix()
	claimsMap["sub"] = subject
	claimsMap["iss"] = appUrl
	claimsMap["aud"] = clientUrlOrIp
	claimsMap["jti"] = uuid.New().String()

	// Set the claims
	token.Claims = jwt.MapClaims(claimsMap)

	// Sign the token
	return token.SignedString([]byte(secret))
}

// Helper function to convert struct to map
func toMap(data interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}
	return result, nil
}

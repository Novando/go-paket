package token

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// Pair holds a raw token and its HMAC-SHA256 hash.
type Pair struct {
	Raw    string
	Hashed string
}

// randomBase64URL generates a URL-safe base64 encoded random byte slice.
func randomBase64URL(n int) (string, error) {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", fmt.Errorf("failed to generate random token: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// hmacSHA256 signs the raw token using HMAC-SHA256 and returns the URL-safe base64 digest.
func hmacSHA256(secret, raw string) string {
	h := hmac.New(sha256.New, []byte(secret))
	// h.Write never returns an error for standard hash implementations.
	_, _ = h.Write([]byte(raw))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// GenerateRefreshTokenPair creates a raw refresh token and its hashed version.
func GenerateRefreshTokenPair(secret string) (Pair, error) {
	raw, err := randomBase64URL(64)
	if err != nil {
		return Pair{}, err
	}
	return Pair{Raw: raw, Hashed: hmacSHA256(secret, raw)}, nil
}

// EncodeRefreshToken returns the HMAC-SHA256 hash of a raw refresh token.
func EncodeRefreshToken(raw, secret string) string {
	return hmacSHA256(secret, raw)
}

// GeneratePasswordResetTokenPair creates a raw password reset token and its hashed version.
func GeneratePasswordResetTokenPair(secret string) (Pair, error) {
	raw, err := randomBase64URL(64)
	if err != nil {
		return Pair{}, err
	}
	return Pair{Raw: raw, Hashed: hmacSHA256(secret, raw)}, nil
}

// EncodePasswordResetToken returns the HMAC-SHA256 hash of a raw password reset token.
func EncodePasswordResetToken(raw, secret string) string {
	return hmacSHA256(secret, raw)
}

// GenerateReactivationTokenPair creates a raw reactivation token and its hashed version.
func GenerateReactivationTokenPair(secret string) (Pair, error) {
	raw, err := randomBase64URL(64)
	if err != nil {
		return Pair{}, err
	}
	return Pair{Raw: raw, Hashed: hmacSHA256(secret, raw)}, nil
}

// EncodeReactivationToken returns the HMAC-SHA256 hash of a raw reactivation token.
func EncodeReactivationToken(raw, secret string) string {
	return hmacSHA256(secret, raw)
}

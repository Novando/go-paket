package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IDTokenBase struct {
	Issuer         string    `json:"iss"`
	Subject        string    `json:"sub"`
	Audience       string    `json:"aud"`
	ExpirationTime time.Time `json:"exp"`
	NotBefore      time.Time `json:"nbf"`
	IssuedAt       time.Time `json:"iat"`
	JwtID          string    `json:"jti"`
}

type Option struct {
	Method   jwt.SigningMethod
	Duration time.Duration
}

var (
	ErrTokenInvalid = errors.New("token is invalid")
)

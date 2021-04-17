package model

import "github.com/dgrijalva/jwt-go"

// AuthTokenClaim ...
type AuthTokenClaim struct {
	*jwt.StandardClaims
	ServicePayload
}

// ServicePayload ...
type ServicePayload struct {
	UserID string `json:"user_id"`
}

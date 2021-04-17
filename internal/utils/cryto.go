package utils

import (
	"errors"
	"fmt"
	"gateway-golang/internal/model"
	"gateway-golang/internal/utils/convert"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SignKey ...
const (
	SignKey = "SignKeyForDevelopingMyService"
)

// GenerateToken ...
func GenerateToken(payload model.ServicePayload, expireAt int64) string {
	claim := model.AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
		ServicePayload: payload,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	ss, _ := token.SignedString([]byte(SignKey))
	return ss
}

// ValidateToken ...
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return nil, fmt.Errorf("claim is invalid")
		}

		now := time.Now().Unix()

		if !claim.VerifyExpiresAt(now, false) {
			return nil, fmt.Errorf("time expires")
		}

		return []byte(SignKey), nil
	})
}

// DecodeToken ...
func DecodeToken(encodedToken string) (*model.AuthTokenClaim, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return "", fmt.Errorf("invalid token %v", token.Header["alg"])
		}
		return []byte(SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claim is invalid")
	}

	authTokenClaim := &model.AuthTokenClaim{}
	convert.ConvertStructToStruct(claims, authTokenClaim)

	return authTokenClaim, nil
}

package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaim struct {
	Id        int64  `json:"id"`
	IsAdmin   bool   `json:"is_admin"`
	ExpiresAt int64  `json:"expires_at"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

const (
	AdminRole    = "admin"
	StandardRole = "standard"
)

func CreateToken(userID int64, isAdmin bool, signingKey string) (string, error) {

	var role string

	if isAdmin {
		role = AdminRole
	} else {
		role = StandardRole
	}
	claims := jwtCustomClaim{
		Id:        userID,
		IsAdmin:   isAdmin,
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		Role:      role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}
	return string(tokenString), nil
}

func VerifyToken(tokenString, signingKey string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid authentication token")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, "", fmt.Errorf("invalid authentication token")
	}

	claims, ok := token.Claims.(*jwtCustomClaim)

	if !ok {
		return 0, "", fmt.Errorf("invalid authentication token")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return 0, "", fmt.Errorf("token has expired")
	}

	return claims.Id, claims.Role, nil
}

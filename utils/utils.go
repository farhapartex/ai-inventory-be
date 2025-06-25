package utils

import (
	"errors"
	"time"

	"github.com/farhapartex/ainventory/config"
	"github.com/farhapartex/ainventory/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(u models.User, TokenType string) (string, error) {
	experiationTime := time.Now().Add(24 * time.Hour)
	claims := config.JWTClaims{
		UserID:    u.ID,
		Email:     u.Email,
		TokenType: TokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(experiationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.JWTSecret)

}

func ValidateJWT(tokenString string) (*config.JWTClaims, error) {
	claims := &config.JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexected signing method")
		}

		return config.JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("Token expired")
	}

	return claims, nil
}

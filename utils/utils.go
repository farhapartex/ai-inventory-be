package utils

import (
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

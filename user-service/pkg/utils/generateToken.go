package utils

import (
	"time"

	"github.com/aurindo10/config"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (*string, error) {
	config.LoadEnv()
	secretKey := config.GetEnv("SECRET_KEY", "JSHKJDSHJKDHSJKDHSJKHDJKSHDJKS")
	var jwtKey = []byte(secretKey)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

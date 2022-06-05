package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/yamess/go-grpc/configs"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), configs.HashingCost)
	return string(bytes), err
}

func VerifyPassword(hashPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	return err == nil
}

func GenerateToken(userId string, email string) (string, error) {
	expirationTime := time.Now().Add(configs.ExpiredTime * time.Minute)
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"alg":     configs.Algorithm,
		"typ":     "JWT",
		"exp":     expirationTime.Unix(),
		"nbf":     time.Now(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(configs.SecretKey))
	return token, err
}

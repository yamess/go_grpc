package utils

import (
	"github.com/yamess/go-grpc/configs"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), configs.HashingCost)
	return string(bytes), err
}

func VerifyPassword(hashPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	return err == nil
}

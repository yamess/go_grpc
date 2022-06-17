package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Host string
var DbUrl string

const (
	HashingCost = 10
	DefaultUser = "32bf3358-b8ae-4d75-bc08-fc4f34d808c8"
	ExpiredTime = 10
	Algorithm   = "HS256"
	SecretKey   = "9ec222277a2fb7a0de4bd59eed949c2afeec427c629dffed20d7b937a014b704"
)

func InitEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Error while loading the environment file")
	}
	Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	DbUrl = os.Getenv("DB_URL")
}

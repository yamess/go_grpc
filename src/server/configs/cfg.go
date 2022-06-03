package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Host string
var DbUrl string

const HashingCost = 10

func InitEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Error while loading the environment file")
	}
	Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	DbUrl = os.Getenv("DB_URL")
}

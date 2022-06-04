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
const DefaultUser = "32bf3358-b8ae-4d75-bc08-fc4f34d808c8"

func InitEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Error while loading the environment file")
	}
	Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	DbUrl = os.Getenv("DB_URL")
}

package db

import (
	"github.com/yamess/go-grpc/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type myDB struct {
	Conn *gorm.DB
}

var MyDB = myDB{}

func (pg *myDB) Connect() {
	db, err := gorm.Open(postgres.Open(configs.DbUrl), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Println(err.Error())
		panic("Failed to connect to database")
	}

	pg.Conn = db
}

func Automigrate(models ...interface{}) {
	MyDB.Connect()

	for _, v := range models {
		ok := MyDB.Conn.AutoMigrate(&v)
		if ok != nil {
			log.Println(ok.Error())
			panic("Failed to apply migration")
		}
	}
}

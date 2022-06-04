package model

import (
	"github.com/google/uuid"
	"github.com/yamess/go-grpc/db"
	"github.com/yamess/go-grpc/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type User struct {
	Id       string `gorm:"primaryKey:unique"`
	Email    string `validate:"email" gorm:"unique"`
	Password string
	IsActive bool
	IsAdmin  bool
	Base
}

// CreateRecord function to create new user
func (user *User) CreateRecord() *gorm.DB {

	// Hash user password
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatalf(err.Error())
	}

	user.Password = hashedPwd
	user.Id = uuid.New().String()
	user.CreatedBy = user.Id
	user.CreatedAt = time.Now().UTC()

	res := db.MyDB.Conn.
		Model(&user).
		Clauses(clause.Returning{}).
		Create(&user)
	return res
}

// GetUser function to a get a single user
func (user *User) GetUser() *gorm.DB {
	res := db.MyDB.Conn.Find(&user)
	return res
}

// UpdateUser function to update an existing user
func (user *User) UpdateUser() *gorm.DB {

	user.UpdatedAt.Time = time.Now().UTC()
	user.UpdatedBy = user.Id

	res := db.MyDB.Conn.
		Model(&user).
		Clauses(clause.Returning{}).
		Omit("Id", "CreatedAt", "CreatedBy", "Password").
		Updates(&user)
	return res
}

// DeleteUser function to delete an existing user
func (user *User) DeleteUser() *gorm.DB {
	res := db.MyDB.Conn.Delete(&user)
	return res
}

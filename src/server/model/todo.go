package model

import (
	"github.com/yamess/go-grpc/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type TodoStatus int

// One Should create the enum manually in the database
const (
	UNKNOWN_STATUS TodoStatus = iota
	PENDING
	STARTED
	COMPLETED
)

type Todo struct {
	Id         uint32 `gorm:"primaryKey"`
	UserId     string
	Text       string
	TodoStatus TodoStatus
	Base
}

func (t *Todo) CreatedTodo() *gorm.DB {
	t.CreatedAt = time.Now().UTC()
	res := db.MyDB.Conn.
		Model(&t).
		Clauses(clause.Returning{}).
		Create(&t)
	return res
}

func (t *Todo) GetTodoById() *gorm.DB {
	res := db.MyDB.Conn.Find(&t)
	return res
}

func (t *Todo) GetTodoList() *gorm.DB {
	res := db.MyDB.Conn.Find(&t)
	return res
}

func (t *Todo) UpdateTodo() *gorm.DB {
	t.UpdatedAt.Time = time.Now().UTC()

	res := db.MyDB.Conn.
		Model(&t).
		Clauses(clause.Returning{}).
		Omit("Id", "UserId", "CreatedAt", "CreatedBy").
		Updates(&t)
	return res
}

func (t *Todo) DeleteTodo() *gorm.DB {
	res := db.MyDB.Conn.Delete(&t)
	return res
}

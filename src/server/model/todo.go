package model

import (
	"github.com/yamess/go-grpc/db"
	pb "github.com/yamess/go-grpc/protos/todo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type TodoStatus int

// One Should create the enum manually in the database
const (
	UNKNOWN_STATUS TodoStatus = iota
	NOT_STARTED
	STARTED
	COMPLETED
)

type Todo struct {
	Id         uint32 `gorm:"primaryKey"`
	UserId     string
	Title      string
	Text       string
	Duration   time.Duration
	StartTime  time.Time
	TodoStatus pb.Status
	Base
}

type TodoList []Todo

func (t *Todo) CreatedTodo() *gorm.DB {
	t.CreatedAt = time.Now()

	res := db.MyDB.Conn.
		Model(&t).
		Clauses(clause.Returning{}).
		Create(&t)
	return res
}

func (t *Todo) GetTodoById() *gorm.DB {
	res := db.MyDB.Conn.Where("Id = ? AND user_id = ?", t.Id, t.UserId).First(&t)
	return res
}

func (tl *TodoList) GetTodoList(userId string) *gorm.DB {
	res := db.MyDB.Conn.Model(Todo{UserId: userId}).Find(&tl)
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

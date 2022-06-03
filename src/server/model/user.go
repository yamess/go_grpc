package model

import (
	"github.com/google/uuid"
	"github.com/yamess/go-grpc/db"
	pb "github.com/yamess/go-grpc/user"
	"github.com/yamess/go-grpc/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type User struct {
	Id        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	IsActive  bool
	IsAdmin   bool
	CreatedAt time.Time
}

type Todo struct {
	Id   string
	Text string
}

func (user *User) GetRPCModel() pb.UserResponse {
	return pb.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsActive:  user.IsActive,
		IsAdmin:   user.IsAdmin,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func (user *User) FromRPC(userGrpc *pb.UserRequest) {
	user.FirstName = userGrpc.FirstName
	user.LastName = userGrpc.LastName
	user.Email = userGrpc.Email
	user.IsAdmin = userGrpc.IsAdmin
	user.IsActive = userGrpc.IsActive
}

// CreateRecord CRUD
func (user *User) CreateRecord() *gorm.DB {
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatalf(err.Error())
	}
	user.Password = hashedPwd
	user.Id = uuid.New().String()
	user.CreatedAt = time.Now()

	res := db.MyDB.Conn.
		Model(&user).
		Clauses(clause.Returning{}).
		Create(&user)
	return res
}

package routes

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/yamess/go-grpc/user"
)

type UserServer struct {
	pb.UserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	var user pb.UserResponse

	id := uuid.New()
	user.Id = id.String()
	return &user, nil
}

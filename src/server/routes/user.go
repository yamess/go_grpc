package routes

import (
	"context"
	"github.com/yamess/go-grpc/model"
	pb "github.com/yamess/go-grpc/user"
)

type UserServer struct {
	pb.UserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	var user model.User
	var userResponse pb.UserResponse

	// Convert Proto buffer data to desired struct
	user.FromRPC(in)

	// Save data in database
	user.CreateRecord()

	// Convert back to proto buffer data type
	userResponse = user.GetRPCModel()

	return &userResponse, nil
}

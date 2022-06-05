package routes

import (
	"context"
	"fmt"
	"github.com/yamess/go-grpc/model"
	pb "github.com/yamess/go-grpc/protos/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type UserServer struct {
	pb.UserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var userResponse pb.CreateUserResponse

	// Convert Proto buffer data to desired struct
	us := model.User{Email: in.Email, Password: in.Password, IsActive: &in.IsActive, IsAdmin: &in.IsAdmin}

	// Save data in database
	res := us.CreateRecord()
	if res.Error != nil {
		log.Printf(res.Error.Error())
		err := status.Error(codes.Internal, res.Error.Error())
		return nil, err
	}

	// Convert back to proto buffer data type
	userResponse = pb.CreateUserResponse{
		Id: us.Id, Email: us.Email, IsActive: *us.IsActive, IsAdmin: *us.IsAdmin,
		CreatedAt: timestamppb.New(us.CreatedAt),
	}
	return &userResponse, nil
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var response pb.GetUserResponse

	// Get id from proto buffer model
	us := model.User{Id: in.Id}

	// Get user from database
	res := us.GetUser()
	if res.Error != nil {
		err := status.Error(codes.NotFound, "id was not found")
		log.Printf("Failed to get user: %s", err)
		return nil, err
	}

	// Convert back to proto buffer data type
	response = pb.GetUserResponse{
		Id: us.Id, Email: us.Email, IsActive: *us.IsActive, IsAdmin: *us.IsAdmin,
		CreatedAt: timestamppb.New(us.CreatedAt),
	}

	return &response, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var response pb.UpdateUserResponse

	us := model.User{
		Id: in.Id, Email: in.Email, IsActive: &in.IsActive, IsAdmin: &in.IsAdmin,
	}
	res := us.UpdateUser()
	if res.Error != nil {
		err := status.Error(codes.Internal, "Unable to updated user")
		log.Printf("Failed to get user: %s", err)
		return nil, err
	}

	response = pb.UpdateUserResponse{
		Id: us.Id, Email: us.Email, IsActive: *us.IsActive, IsAdmin: *us.IsAdmin,
		CreatedAt: timestamppb.New(us.CreatedAt), UpdatedAt: timestamppb.New(us.UpdatedAt.Time),
	}

	return &response, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	us := model.User{Id: in.Id}

	res := us.DeleteUser()
	if res.Error != nil {
		err := status.Error(codes.Internal, "Unable to delete user")
		log.Printf("Failed to delete user: %s", res.Error)
		return nil, err
	}
	responseMessage := fmt.Sprintf("User %s deleted with success", us.Id)
	response := pb.DeleteUserResponse{Message: responseMessage}

	return &response, nil
}

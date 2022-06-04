package routes

import (
	"context"
	"github.com/yamess/go-grpc/configs"
	"github.com/yamess/go-grpc/model"
	pb "github.com/yamess/go-grpc/protos/todo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type TodoServer struct {
	pb.TodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {

	todo := model.Todo{UserId: in.UserId, Text: in.Text}
	todo.CreatedBy = configs.DefaultUser
	todo.CreatedAt = time.Now().UTC()

	res := todo.CreatedTodo()
	if res.Error != nil {
		log.Printf("Unable to create todo: %s", res.Error)
		err := status.Error(codes.Internal, res.Error.Error())
		return nil, err
	}
	response := pb.CreateTodoResponse{
		Id: todo.Id, UserId: todo.UserId, Text: todo.Text, Status: pb.Status(todo.TodoStatus),
		CreatedAt: timestamppb.New(todo.CreatedAt),
	}
	return &response, nil
}

func (s *TodoServer) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	todo := model.Todo{Id: in.Id, UserId: in.UserId}
	res := todo.GetTodoById()
	if res.Error != nil {
		log.Printf("Unable to read todo: %s", res.Error)
		err := status.Error(codes.Internal, res.Error.Error())
		return nil, err
	}
	response := pb.GetTodoResponse{
		Id: todo.Id, UserId: todo.UserId, Text: todo.Text, Status: pb.Status(todo.TodoStatus),
		CreatedAt: timestamppb.New(todo.CreatedAt),
	}
	return &response, nil
}

package routes

import (
	"context"
	"errors"
	"fmt"
	"github.com/yamess/go-grpc/configs"
	"github.com/yamess/go-grpc/model"
	pb "github.com/yamess/go-grpc/protos/todo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"log"
	"time"
)

type TodoServer struct {
	pb.TodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {

	todo := model.Todo{
		UserId:     in.UserId,
		Title:      in.Title,
		Text:       in.Text,
		Duration:   in.Duration.AsDuration(),
		StartTime:  in.StartTime.AsTime(),
		TodoStatus: pb.Status_NOT_STARTED,
	}
	todo.CreatedAt = time.Now().UTC()

	if todo.StartTime.Before(time.Now()) {
		todo.TodoStatus = pb.Status_STARTED
	}

	todo.CreatedBy = configs.DefaultUser // To be replaced @TODO

	res := todo.CreatedTodo()
	if res.Error != nil {
		log.Printf("Unable to create todo: %s", res.Error)
		err := status.Error(codes.Internal, res.Error.Error())
		return nil, err
	}
	response := pb.CreateTodoResponse{
		Id:        todo.Id,
		UserId:    todo.UserId,
		Title:     todo.Title,
		Text:      todo.Text,
		Duration:  durationpb.New(todo.Duration),
		StartTime: timestamppb.New(todo.StartTime),
		Status:    todo.TodoStatus,
		CreatedAt: timestamppb.New(todo.CreatedAt),
	}
	return &response, nil
}

func (s *TodoServer) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	todo := model.Todo{Id: in.Id, UserId: in.UserId}
	res := todo.GetTodoById()

	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		log.Printf("Unable to read todo: %s", res.Error)
		err := status.Error(codes.Internal, res.Error.Error())
		return nil, err
	} else if res.RowsAffected == 0 {
		log.Printf("No record found for this user id")
		err := status.Error(codes.NotFound, "No record found for this user id")
		return nil, err
	}

	response := pb.GetTodoResponse{
		Id:        todo.Id,
		UserId:    todo.UserId,
		Title:     todo.Title,
		Text:      todo.Text,
		Duration:  durationpb.New(todo.Duration),
		StartTime: timestamppb.New(todo.StartTime),
		Status:    todo.TodoStatus,
		CreatedAt: timestamppb.New(todo.CreatedAt),
	}
	return &response, nil
}

func (s *TodoServer) GetTodoList(in *pb.GetTodoListRequest, stream pb.TodoService_GetTodoListServer) error {
	var todoList model.TodoList

	res := todoList.GetTodoList(in.UserId)
	if res.Error != nil {
		log.Printf("Unable to read todo list: %s", res.Error)
		err := status.Error(codes.Internal, res.Error.Error())
		return err
	}

	for _, todo := range todoList {
		response := pb.GetTodoResponse{
			Id:        todo.Id,
			UserId:    todo.UserId,
			Title:     todo.Title,
			Text:      todo.Text,
			Duration:  durationpb.New(todo.Duration),
			StartTime: timestamppb.New(todo.StartTime),
			Status:    todo.TodoStatus,
			CreatedAt: timestamppb.New(todo.CreatedAt),
		}
		if err := stream.Send(&response); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *TodoServer) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	todo := model.Todo{
		Id:         in.Id,
		UserId:     in.UserId,
		Title:      in.Title,
		Text:       in.Text,
		Duration:   in.Duration.AsDuration(),
		StartTime:  in.StartTime.AsTime(),
		TodoStatus: in.Status,
	}

	todo.UpdatedBy = configs.DefaultUser
	todo.UpdatedAt.Time = time.Now().UTC()

	res := todo.UpdateTodo()
	if res.Error != nil {
		log.Printf("Unable to update todo: %s", res.Error)
		err := status.Error(codes.Internal, res.Error.Error())
		return nil, err
	}
	response := pb.UpdateTodoResponse{
		Id:        todo.Id,
		UserId:    todo.UserId,
		Title:     todo.Title,
		Text:      todo.Text,
		Status:    todo.TodoStatus,
		UpdatedAt: timestamppb.New(todo.UpdatedAt.Time),
		CreatedAt: timestamppb.New(todo.CreatedAt),
	}
	return &response, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	todo := model.Todo{Id: in.Id, UserId: in.UserId}

	res := todo.DeleteTodo()
	if res.Error != nil {
		log.Printf("Unable to delete todo: %s", res.Error)
		err := status.Error(codes.Internal, res.Error.Error())
		return nil, err
	}
	responseMessage := fmt.Sprintf("Todo with id %d deleted with success", todo.Id)
	response := pb.DeleteTodoResponse{Message: responseMessage}
	return &response, nil
}

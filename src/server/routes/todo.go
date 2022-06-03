package routes

import (
	"context"
	pb "github.com/yamess/go-grpc/todo"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
)

type TodoServer struct {
	pb.TodoServiceServer
}

// TodoDB Mimic a database
var TodoDB pb.TodoItems

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.TodoItem) (*pb.TodoItem, error) {
	var item pb.TodoItem

	item = pb.TodoItem{Text: in.GetText()}
	item.Id = rand.Int31()

	TodoDB.Items = append(TodoDB.Items, &item)

	return &item, nil
}

func (s *TodoServer) ReadTodo(ctx context.Context, empty *emptypb.Empty) (*pb.TodoItems, error) {
	return &TodoDB, nil
}

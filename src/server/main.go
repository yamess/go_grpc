package main

import (
	"github.com/yamess/go-grpc/configs"
	"github.com/yamess/go-grpc/db"
	"github.com/yamess/go-grpc/interceptors"
	"github.com/yamess/go-grpc/model"
	tpb "github.com/yamess/go-grpc/protos/todo"
	upb "github.com/yamess/go-grpc/protos/user"
	"github.com/yamess/go-grpc/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	configs.InitEnv()

	db.Automigrate(model.Todo{}, model.User{})

	lis, err := net.Listen("tcp", configs.Host)

	if err != nil {
		log.Fatalf(err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.UnaryInterceptors),
		grpc.StreamInterceptor(interceptors.StreamInterceptors),
	)

	tpb.RegisterTodoServiceServer(grpcServer, &routes.TodoServer{})
	upb.RegisterUserServiceServer(grpcServer, &routes.UserServer{})
	reflection.Register(grpcServer)

	log.Printf("myserver listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}

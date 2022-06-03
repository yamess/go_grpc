package main

import (
	"github.com/yamess/go-grpc/interceptors"
	"github.com/yamess/go-grpc/routes"
	tpb "github.com/yamess/go-grpc/todo"
	upb "github.com/yamess/go-grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

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

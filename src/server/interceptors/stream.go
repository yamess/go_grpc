package interceptors

import (
	"google.golang.org/grpc"
	"log"
)

func StreamInterceptors(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

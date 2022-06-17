package main

import (
	"github.com/gin-gonic/gin"
	proto "github.com/yamess/go-grpc/client/protos/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:9000", opts)
	if err != nil {
		panic(err.Error())
	}

	client := proto.NewUserServiceClient(conn)
	//md := metadata.New(map[string]string{"Authorization": "Bearer 232dsfdfd"})

	g := gin.Default()

	g.POST("/create/user", func(ctx *gin.Context) {
		var user proto.CreateUserRequest
		err = ctx.BindJSON(&user)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if response, err := client.CreateUser(ctx, &user); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": response})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run rest api myserver: %s", err.Error())
	}

}

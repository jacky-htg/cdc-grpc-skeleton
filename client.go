package main

import (
	"fmt"
	"skeleton/pb/users"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7070", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %s", err)
		return
	}
	defer conn.Close()

	ctx := context.Background()

	user := users.NewUserServiceClient(conn)

	response, err := user.List(ctx, &users.UserListInput{})

	if err != nil {
		fmt.Printf("Error when calling grpc service: %s", err)
		return
	}
	fmt.Printf("Resp received: %v", response)
}

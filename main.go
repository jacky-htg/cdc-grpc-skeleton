package main

import (
	"context"
	"log"
	"net"
	"os"
	"skeleton/pb/users"

	"google.golang.org/grpc"
)

func main() {

	log := log.New(os.Stdout, "grpc skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// listen tcp port
	lis, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()

	// routing grpc services
	grpcRoute(grpcServer, log)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
	log.Print("serve grpc on port: 7070")

}

func grpcRoute(grpcServer *grpc.Server, log *log.Logger) {
	handler := new(userHandler)
	handler.log = log
	users.RegisterUserServiceServer(grpcServer, handler)
}

type userHandler struct {
	log *log.Logger
}

func (u *userHandler) List(ctx context.Context, in *users.UserListInput) (*users.Users, error) {
	var list users.Users
	list.User = append(list.User, &users.User{Id: 1, Username: "Jacky"})
	list.User = append(list.User, &users.User{Id: 2, Username: "Jet Lee", Email: "jet@lee.com"})
	return &list, nil
}

func (u *userHandler) Create(ctx context.Context, in *users.UserCreateInput) (*users.User, error) {
	return &users.User{}, nil
}

package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/turao/go-cards/users"
	pb "github.com/turao/go-cards/users/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUsersServer

	createUser users.CreateUserUseCase
	getUser    users.GetUserUseCase
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	log.Println("[server] Getting user...")

	uuid, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, errors.New("unable to parse user id")
	}

	found, err := s.getUser.Execute(uuid)
	if err != nil {
		return nil, err // todo: avoid naked errors!
	}

	return &pb.User{Id: found.Id().String(), Name: found.Name()}, nil
}

// func (s *server) ListUsers(*pb.ListUsersRequest, Users_ListUsersServer) error {
// 	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
// }

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	log.Println("[server] Creating user...")

	user, err := s.createUser.Execute(req.GetName())
	if err != nil {
		return nil, errors.New("unable to create new user")
	}
	return &pb.User{Id: user.Id().String(), Name: user.Name()}, nil
}

const port = ":8080"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Failed to listen")
	}
	s := grpc.NewServer()

	repository := users.NewInMemoryRepository()

	pb.RegisterUsersServer(s, &server{
		createUser: users.CreateUserUseCase{Repository: repository},
		getUser:    users.GetUserUseCase{Repository: repository},
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve")
	}
}

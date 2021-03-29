package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/turao/go-cards/users"
	pb "github.com/turao/go-cards/users/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUsersServer

	repository *users.Repository
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	fmt.Println("Getting user...")
	uuid, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errors.New("unable to parse user id")
	}

	user, err := s.repository.GetUser(uuid)
	if err != nil {
		return nil, err // todo: avoid naked errors!
	}

	return &pb.User{Id: user.Id().String(), Name: user.Name()}, nil
}

// func (s *server) ListUsers(*pb.ListUsersRequest, Users_ListUsersServer) error {
// 	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
// }

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := users.New(req.GetName())

	user, err := s.repository.CreateUser(user)
	if err != nil {
		return nil, err // todo: avoid naked errors
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

	pb.RegisterUsersServer(s, &server{repository: users.NewRepository()})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve")
	}
}

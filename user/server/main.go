package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/turao/go-cards/user"
	pb "github.com/turao/go-cards/user/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUsersServer

	getUser    user.GetUserUseCase
	listUsers  user.ListUsersUseCase
	createUser user.CreateUserUseCase
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

func (s *server) ListUsers(req *pb.ListUsersRequest, stream pb.Users_ListUsersServer) error {
	users, err := s.listUsers.Execute()
	if err != nil {
		return err // todo: avoid naked errors!
	}

	for _, user := range users {
		err = stream.Send(&pb.User{Id: user.Id().String(), Name: user.Name()})
		if err != nil {
			return err
		}
	}

	return nil
}

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

	repository := user.NewInMemoryRepository()

	pb.RegisterUsersServer(s, &server{
		getUser:    user.GetUserUseCase{Repository: repository},
		listUsers:  user.ListUsersUseCase{Repository: repository},
		createUser: user.CreateUserUseCase{Repository: repository},
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve")
	}
}

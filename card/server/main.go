package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/turao/go-cards/card"
	pb "github.com/turao/go-cards/card/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCardsServer

	createCard    card.CreateCardUseCase
	assignOwner   card.AssignOwnerUseCase
	unassignOwner card.UnassignOwnerUseCase
}

func (s *server) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.Card, error) {
	log.Println("[server]", "creating card...")

	card, err := s.createCard.Execute()
	if err != nil {
		return nil, errors.New("unable to create new card")
	}
	return &pb.Card{Id: card.Id().String(), OwnerId: card.OwnerId().String()}, nil
}

func (s *server) AssignOwner(ctx context.Context, req *pb.AssignOwnerRequest) (*pb.Card, error) {
	log.Println("[server]", "assigning owner to card...")
	cardId, err := uuid.Parse(req.CardId)
	if err != nil {
		return nil, errors.New("unable to parse cardId")
	}

	ownerId, err := uuid.Parse(req.OwnerId)
	if err != nil {
		return nil, errors.New("unable to parse ownerId")
	}

	card, err := s.assignOwner.Execute(cardId, ownerId)
	if err != nil {
		return nil, errors.New("failed to add card to user")
	}

	return &pb.Card{Id: card.Id().String(), OwnerId: card.OwnerId().String()}, nil
}

const port = ":8081"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen")
	}
	s := grpc.NewServer()

	repository := card.NewInMemoryRepository()

	pb.RegisterCardsServer(s, &server{
		createCard:    card.CreateCardUseCase{Repository: repository},
		assignOwner:   card.AssignOwnerUseCase{Repository: repository},
		unassignOwner: card.UnassignOwnerUseCase{Repository: repository},
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve")
	}
}

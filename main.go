package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	cardGRPC "github.com/turao/go-cards/card/grpc"
	userGRPC "github.com/turao/go-cards/user/grpc"
)

func PrettyPrintln(data json.Marshaler) {
	d, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	log.Println("[main]", string(d))
}

func main() {
	connUser, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Unable to connect (grpc): ", err.Error())
	}
	defer connUser.Close()

	userClient := userGRPC.NewUsersClient(connUser)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := userClient.CreateUser(ctx, &userGRPC.CreateUserRequest{})
	if err != nil {
		log.Println("[main]", "Unable to create user: ", err.Error())
	}
	log.Println("[main]", "Created User...")
	log.Println("[main]", user)

	user, err = userClient.GetUser(ctx, &userGRPC.GetUserRequest{UserId: user.GetId()})
	if err != nil {
		log.Println("[main]", "Unable to get user: ", err.Error())
	}
	log.Println("[main]", "Got User...", user)

	users, err := userClient.ListUsers(ctx, &userGRPC.ListUsersRequest{})
	if err != nil {
		log.Println("[main]", "Unable to list users: ", err.Error())
	}

	log.Println("[main]", "Listing all users now...")
	for {
		u, err := users.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return
		}
		log.Println("[main]", "Got User...", u)
	}

	// Create a new card
	connCard, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Unable to connect (grpc): ", err.Error())
	}
	defer connCard.Close()

	cardClient := cardGRPC.NewCardsClient(connCard)
	card, err := cardClient.CreateCard(ctx, &cardGRPC.CreateCardRequest{})
	if err != nil {
		log.Println("[main]", "unable to create card", err.Error())
	}

	// Add this card to the user
	_, err = userClient.AddCard(ctx, &userGRPC.AddCardRequest{UserId: user.GetId(), CardId: card.GetId()})
	if err != nil {
		log.Println("[main]", "unable to add card to user (first time)", err.Error())
	}

	_, err = userClient.AddCard(ctx, &userGRPC.AddCardRequest{UserId: user.GetId(), CardId: card.GetId()})
	if err != nil {
		log.Println("[main]", "unable to add card to user (second time)", err.Error())
	}

	// assign user as owner of this card
	card, err = cardClient.AssignOwner(ctx, &cardGRPC.AssignOwnerRequest{CardId: card.GetId(), OwnerId: user.GetId()})
	if err != nil {
		log.Println("[main]", "unable to assign owner to card", err.Error())
	}
	log.Println("[main]", "owner assigned to card", card)
}

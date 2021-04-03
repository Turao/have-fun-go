package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/turao/go-cards/user/grpc"
)

func PrettyPrintln(data json.Marshaler) {
	d, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	log.Println("[main]", string(d))
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Unable to connect (grpc): ", err.Error())
	}
	defer conn.Close()

	client := pb.NewUsersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	u, err := client.CreateUser(ctx, &pb.CreateUserRequest{})
	if err != nil {
		log.Println("[main]", "Unable to create user: ", err.Error())
	}
	log.Println("[main]", "Created User...")
	log.Println("[main]", u)

	u, err = client.GetUser(ctx, &pb.GetUserRequest{UserId: u.Id})
	if err != nil {
		log.Println("[main]", "Unable to get user: ", err.Error())
	}
	log.Println("[main]", "Got User...", u)

	users, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
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

	cardId := "00000000-0000-0000-0000-000000000000"
	_, err = client.AddCard(ctx, &pb.AddCardRequest{UserId: u.Id, CardId: cardId})
	if err != nil {
		log.Println("[main]", "unable to add card to user (first time)", err.Error())
	}

	_, err = client.AddCard(ctx, &pb.AddCardRequest{UserId: u.Id, CardId: cardId})
	if err != nil {
		log.Println("[main]", "unable to add card to user (second time)", err.Error())
	}

}

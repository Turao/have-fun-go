package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/turao/go-cards/auctions"
	"github.com/turao/go-cards/cards"
	"github.com/turao/go-cards/users"
	"google.golang.org/grpc"

	pb "github.com/turao/go-cards/users/grpc"
)

func PrettyPrintln(data json.Marshaler) {
	d, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	log.Println("[main]", string(d))
}

func main() {

	user := users.New("joe")
	PrettyPrintln(user)

	card := cards.New()
	_, err := card.AssignOwner(user.Id())
	if err != nil {
		panic(err.Error())
	}

	_, err = user.AddCard(card.Id())
	if err != nil {
		panic(err.Error())
	}

	PrettyPrintln(user)

	auction := auctions.New()
	PrettyPrintln(auction)

	_, err = auction.Start()
	if err != nil {
		panic(err.Error())
	}
	PrettyPrintln(auction)

	_, err = auction.PlaceBid(user.Id(), card.Id(), 10)
	if err != nil {
		panic(err.Error())
	}

	_, err = auction.End()
	if err != nil {
		panic(err.Error())
	}

	PrettyPrintln(auction)

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

}

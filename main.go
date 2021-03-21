package main

import (
	"encoding/json"
	"fmt"

	"github.com/turao/go-cards/auctions"
	"github.com/turao/go-cards/cards"
	"github.com/turao/go-cards/users"
)

func PrettyPrintln(data json.Marshaler) {
	d, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(d))
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

}

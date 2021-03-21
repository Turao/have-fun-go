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
	card.AssignOwner(user.Id())
	user.AddCard(card.Id())
	PrettyPrintln(user)

	auction := auctions.New()
	PrettyPrintln(auction)
	auction.Start()
	PrettyPrintln(auction)

	auction.PlaceBid(user.Id(), card.Id(), 10)

	auction.End()
	PrettyPrintln(auction)

}

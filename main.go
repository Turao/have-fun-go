package main

import (
	"encoding/json"
	"fmt"

	"github.com/turao/go-cards/cards"
	"github.com/turao/go-cards/users"
)

func PrettyPrintln(data interface{}) {
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
	PrettyPrintln(card)
	card.AssignOwner(user.Id())

	user.AddCard(card.Id())
	PrettyPrintln(user)
}

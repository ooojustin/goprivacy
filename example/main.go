package main

import (
	"fmt"

	"github.com/ooojustin/goprivacy"
)

var client goprivacy.Client

func main() {
	client = goprivacy.Client{Key: "bbb0e10d-8aad-4f28-9800-c69e4b0d25cf"}
	//listCards(1)
	//retrieveCard("c35a015e-3cec-460b-92f3-8d19b63f50ac")
	//listFundingAccounts()
}

func listCards(page int) {
	// retrieve a list of cards
	if cards, err := client.ListCards(page); err == nil {
		for _, card := range *cards {
			fmt.Println(card)
		}
	} else {
		fmt.Println(err)
	}
}

func retrieveCard(token string) {
	// retrieve a single card from token
	if card, err := client.GetCard(token); err == nil {
		fmt.Println(card)
	} else {
		fmt.Println(err)
	}
}

func listFundingAccounts() {
	if accounts, err := client.ListFundingAccounts(); err == nil {
		fmt.Println(accounts)
	} else {
		fmt.Println(err)
	}
}

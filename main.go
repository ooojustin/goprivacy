package main

import (
	"fmt"

	"justin.ooo/privacy/api"
)

func main() {

	client := api.PrivacyClient{Key: "bbb0e10d-8aad-4f28-9800-c69e4b0d25cf"}

	// retrieve a list of cards
	//if cards, err := client.GetCards(1); err == nil {
	//for _, card := range *cards {
	//fmt.Println(card)
	//}
	//} else {
	//fmt.Println(err)
	//}

	// retrieve a single card from token
	if card, err := client.GetCard("c35a015e-3cec-460b-92f3-8d19b63f50ac"); err == nil {
		fmt.Println(card)
	} else {
		fmt.Println(err)
	}

}

package goprivacy

import (
	"encoding/json"
	"strconv"
)

type CardResponse struct {
	Data         []Card `json:"data"`
	TotalEntries int    `json:"total_entries"`
	TotalPages   int    `json:"total_pages"`
	Page         int    `json:"page"`
}

type Card struct {
	Created            string `json:"created"`
	Token              string `json:"token"`
	LastFour           string `json:"last_four"`
	Hostname           string `json:"hostname"`
	Mem                string `json:"memo"`
	Type               string `json:"type"`
	SpendLimit         int    `json:"spend_limit"`
	SpendLimitDuration string `json:"spend_limit_duration"`
	State              string `json:"state"`
}

func (pc Client) ListCards(page int) (*[]Card, error) {

	dest := BaseURL + "card?page=" + strconv.Itoa(page)
	body, err := pc.GET(dest)
	if err != nil {
		return nil, err
	}

	var cr CardResponse
	if err = json.Unmarshal(*body, &cr); err != nil {
		return nil, err
	}

	return &cr.Data, nil

}

func (pc Client) GetCard(cardToken string) (*Card, error) {

	dest := BaseURL + "card?card_token=" + cardToken
	body, err := pc.GET(dest)
	if err != nil {
		return nil, err
	}

	var cr CardResponse
	if err = json.Unmarshal(*body, &cr); err != nil {
		return nil, err
	}

	return &cr.Data[0], nil

}

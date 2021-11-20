package api

import (
	"encoding/json"
	"strconv"
)

type CardResponse struct {
	Data         []PrivacyCard `json:"data"`
	TotalEntries int           `json:"total_entries"`
	TotalPages   int           `json:"total_pages"`
	Page         int           `json:"page"`
}

type PrivacyCard struct {
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

func (pc PrivacyClient) GetCards(page int) (*[]PrivacyCard, error) {

	url := BaseURL + "card?page=" + strconv.Itoa(page)
	body, err := pc.GET(url)
	if err != nil {
		return nil, err
	}

	var cr CardResponse
	if err = json.Unmarshal(*body, &cr); err != nil {
		return nil, err
	}

	return &cr.Data, nil

}

func (pc PrivacyClient) GetCard(cardToken string) (*PrivacyCard, error) {

	url := BaseURL + "card?card_token=" + cardToken
	body, err := pc.GET(url)
	if err != nil {
		return nil, err
	}

	var cr CardResponse
	if err = json.Unmarshal(*body, &cr); err != nil {
		return nil, err
	}

	return &cr.Data[0], nil

}

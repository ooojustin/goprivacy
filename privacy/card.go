package goprivacy

import (
	"encoding/json"
	"fmt"
	"net/url"
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
	Memo               string `json:"memo"`
	Type               string `json:"type"`
	SpendLimit         int    `json:"spend_limit"`
	SpendLimitDuration string `json:"spend_limit_duration"`
	State              string `json:"state"`
}

func (pc Client) ListCards(page int, params map[string]string) (*[]Card, error) {

	params["page"] = strconv.Itoa(page)
	uv := url.Values{}
	for k, v := range params {
		uv.Add(k, v)
	}

	dest := BaseURL + "card?" + uv.Encode()
	fmt.Println(dest)
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

	cards, err := pc.ListCards(1, map[string]string{"card_token": cardToken})
	if err != nil {
		return nil, err
	} else {
		return &(*cards)[0], nil
	}

}

package goprivacy

import (
	"encoding/json"
)

type FundingAccount struct {
	Created     string `json:"created"`
	AccountName string `json:"account_name"`
	LastFour    string `json:"last_four"`
	Nickname    string `json:"nickname"`
	State       string `json:"statte"`
	Token       string `json:"token"`
	Type        string `json:"type"`
}

func (pc Client) ListFundingAccounts() (*[]FundingAccount, error) {

	url := BaseURL + "fundingsource"
	body, err := pc.GET(url)
	if err != nil {
		return nil, err
	}

	var accounts []FundingAccount
	if err = json.Unmarshal(*body, &accounts); err != nil {
		return nil, err
	}

	return &accounts, nil

}

package goprivacy

import (
	"encoding/json"
)

type TransactionResponse struct {
	Data         []Transaction `json:"data"`
	TotalEntries int           `json:"total_entries"`
	TotalPages   int           `json:"total_pages"`
	Page         int           `json:"page"`
}

type TFunding struct {
	Amount interface{} `json:"amount"` // has been both int and string
	Token  string      `json:"token"`
	Type   string      `json:"string"`
}

type Transaction struct {
	Amount        int        `json:"amount"`
	Card          Card       `json:"card"`
	Events        []Event    `json:"events"`
	Funding       []TFunding `json:"funding"`
	Merchant      Merchant   `json:"merchant"`
	Result        string     `json:"result"`
	SettledAmount int        `json:"settled_amount"`
	Status        string     `json:"status"`
	Token         string     `json:"token"`
}

func (pc Client) ListTransactions(status string) (*[]Transaction, error) {

	url := BaseURL + "transaction/" + status
	body, err := pc.GET(url)
	if err != nil {
		return nil, err
	}

	var tr TransactionResponse
	if err = json.Unmarshal(*body, &tr); err != nil {
		return nil, err
	}

	return &tr.Data, nil

}

func (pc Client) GetTransaction(transactionToken string) (*Transaction, error) {

	url := BaseURL + "transaction/all?transaction_token=" + transactionToken
	body, err := pc.GET(url)
	if err != nil {
		return nil, err
	}

	var tr TransactionResponse
	if err = json.Unmarshal(*body, &tr); err != nil {
		return nil, err
	}

	return &tr.Data[0], nil

}

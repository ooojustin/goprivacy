package goprivacy

import (
	"encoding/json"
	"net/url"
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

func (pc Client) ListAllTransactions() (*[]Transaction, error) {
	return pc.ListTransactions("all", map[string]string{})
}

func (pc Client) ListTransactions(status string, params map[string]string) (*[]Transaction, error) {

	dest := BaseURL + "transaction/" + status + "?"
	uv := url.Values{}
	for k, v := range params {
		uv.Add(k, v)
	}
	dest += uv.Encode()

	body, err := pc.GET(dest)
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

	dest := BaseURL + "transaction/all?transaction_token=" + transactionToken
	body, err := pc.GET(dest)
	if err != nil {
		return nil, err
	}

	var tr TransactionResponse
	if err = json.Unmarshal(*body, &tr); err != nil {
		return nil, err
	}

	return &tr.Data[0], nil

}

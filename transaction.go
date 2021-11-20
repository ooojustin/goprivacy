package goprivacy

type TFunding struct {
	Amount int `json:"amount"`
	Token string `json:"token"`
	Type string `json:"string"`
}

type Transaction struct {
	Amount int `json:"amount"`
	Card Card `json:"card"`
	Events []Event `json:"events"`
	Funding []TFunding `json:"funding"`
	Merchant Merchant `json:"merchant"`
	Result string `json:"result"`
	SettledAmount int `json:"settled_amount"`
	Status string `json:"status"`
	Token string `json:"token"`
}

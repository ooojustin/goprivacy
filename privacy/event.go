package goprivacy

type Event struct {
	Amount  int    `json:"amount"`
	Created string `json:"created"`
	Result  string `json:"result"`
	Token   string `json:"token"`
	Type    string `json:"type"`
}

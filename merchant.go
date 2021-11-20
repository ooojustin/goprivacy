package goprivacy

type Merchant struct {
	AcceptorID string `json:"acceptor_id"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Descriptor string `json:"descriptor"`
	MCC        string `json:"mcc"`
	State      string `json:"state"`
}

package domain

type Client struct {
	Name     string     `json:"name"`
	Document string     `json:"document"`
	Address  Address    `json:"address"`
	Credit   CreditInfo `json:"credit_status"`
}

type Address struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`
	District string `json:"district"`
}

type CreditInfo struct {
	Situation string `json:"situation"`
	Risk      string `json:"risk"`
	Type      string `json:"type"`
}

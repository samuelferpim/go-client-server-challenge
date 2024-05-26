package model

type ExchangeRate struct {
	Bid   string `json:"bid,omitempty"`
	Error string `json:"error,omitempty"`
}

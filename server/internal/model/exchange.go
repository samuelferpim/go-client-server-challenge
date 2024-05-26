package model

type Conversion struct {
	ExchangeRate ExchangeRate `json:"USDBRL"`
}

type ExchangeRate struct {
	Id  string `json:"-"`
	Bid string `json:"bid"`
}

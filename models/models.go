package models

type Order struct {
	Customer string `json:"customer"`
	Pizza    string `json:"pizza"`
	Quantity int    `json:"quantity"`
}

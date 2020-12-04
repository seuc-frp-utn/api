package payments

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	UUID string `json:"uuid"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
	Payer string `json:"payer"`
}

type PaymentCreate struct {
	Description string `json:"description"`
	Amount float64 `json:"amount"`
	Payer string `json:"payer"`
}

type PaymentUpdate struct {
	Description string `json:"description"`
	Amount float64 `json:"amount"`
	Payer string `json:"payer"`
}
package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}
type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	Lastname  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		FirstName: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(10000000)),
		CreatedAt: time.Now().UTC(),
	}
}

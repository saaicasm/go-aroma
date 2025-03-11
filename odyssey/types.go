package main

import "math/rand"

type Account struct {
	ID        int `json:"id"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Number    int64 `json:"number"`
	Balance   int64 `json:"balance"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(10000000)),
	}
}

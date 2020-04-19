package model

import (
	"errors"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountRepo interface {
	Add(acc Account) (err error)
	Get(id string) (account Account, err error)
	Update(id string, name string) (err error)
}
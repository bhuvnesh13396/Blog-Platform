package account

import (
	"context"
	"errors"

	"sample/model"
)

var (
	errInvalidArgument = errors.New("invalid argument")
	errInvalidId       = errors.New("No such account exists with given Id")
)

type Service interface {
	GetAccount(ctx context.Context, id string) (account model.Account, err error)
	AddAccount(ctx context.Context, id string, name string) (err error)
	UpdateAccount(ctx context.Context, id string, name string) (err error)
}

type service struct {
	accountRepo model.AccountRepo
}

func NewService(accountRepo model.AccountRepo) Service {
	return &service{
		accountRepo: accountRepo,
	}
}

func (s *service) GetAccount(ctx context.Context, id string) (account model.Account, err error) {
	if len(id) < 1 {
		err = errInvalidArgument
		return
	}
	return s.accountRepo.Get(id)
}

func (s *service) AddAccount(ctx context.Context, id string, name string) (err error) {
	if len(id) < 1 || len(name) < 1 {
		err = errInvalidArgument
		return
	}

	acc := model.Account{
		ID:   id,
		Name: name,
	}

	return s.accountRepo.Add(acc)
}

func (s *service) UpdateAccount(ctx context.Context, id string, name string) (err error) {
	if len(id) < 1 || len(name) < 1 {
		err = errInvalidArgument
		return
	}

	return s.accountRepo.Update(id, name)

}
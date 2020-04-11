package repo

import (
	"encoding/json"

	"sample/kv"
	"sample/model"
)

type accountRepo struct {
	db *kv.Database
}

func NewAccountRepo(db *kv.Database) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (repo *accountRepo) Add(a model.Account) (err error) {
	data, err := json.Marshal(a)
	if err != nil {
		return
	}
	err = repo.db.Set(a.ID, data)
	return
}

func (repo *accountRepo) Get(id string) (a model.Account, err error) {
	data, err := repo.db.Get(id)
	if err != nil {
		err = model.ErrAccountNotFound
		return
	}

	err = json.Unmarshal(data, &a)
	if err != nil {
		return
	}

	return
}

func (repo *accountRepo) Update(id string, name string) (err error) {
	a, err := repo.Get(id)
	if err != nil {
		return
	}
	a.Name = name
	data, err := json.Marshal(a)
	if err != nil {
		return
	}
	err = repo.db.Set(a.ID, data)
	return
}

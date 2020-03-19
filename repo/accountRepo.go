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

func (repo *accountRepo) Get(instanceDomain string) (a model.Account, err error) {
	data, err := repo.db.Get(instanceDomain)
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

func (repo *accountRepo) Update(a model.Account, name string) (err error) {
	data, err := json.Marshal(a)
	if err != nil {
		return
	}
	err = repo.db.Set(a.ID, data)
	return
}

package psql

import (
	"database/sql"
	"sample/account/model"
)

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (repo *accountRepo) Add(a model.Account) (err error) {
	q := "INSERT INTO account (id, name) VALUES ($1, $2)"
	_, err = repo.db.Exec(q, a.ID, a.Name)
	return
}

func (repo *accountRepo) Get(id string) (a model.Account, err error) {
	row := repo.db.QueryRow("SELECT * FROM account WHERE id = $1", id)
	switch err := row.Scan(&a.ID, &a.Name); err {
	case sql.ErrNoRows:
		return a, model.ErrAccountNotFound
	case nil:
		return a, nil
	}

	return
}

func (repo *accountRepo) Update(id string, name string) (err error) {
	q := "UPDATE account SET name = $2 WHERE ID = $1"
	_, err = repo.db.Exec(q, id, name)
	return
}
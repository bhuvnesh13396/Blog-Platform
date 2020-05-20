package psql

import (
	"database/sql"

	"sample/model"
)

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) (*accountRepo, error) {
	query := "CREATE TABLE IF NOT EXISTS account (id varchar primary key, name varchar, username varchar, password varchar)"
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &accountRepo{
		db: db,
	}, nil
}

func (repo *accountRepo) Add(a model.Account) (err error) {
	q := "INSERT INTO account (id, name, username, password) VALUES ($1, $2, $3, $4)"
	_, err = repo.db.Exec(q, a.ID, a.Name, a.Username, a.Password)
	return
}

func (repo *accountRepo) Get(id string) (a model.Account, err error) {
	row := repo.db.QueryRow("SELECT * FROM account WHERE id = $1", id)
	err = row.Scan(&a.ID, &a.Name, &a.Username, &a.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			err = model.ErrAccountNotFound
		}
		return
	}

	return
}

func (repo *accountRepo) Get1(username string) (a model.Account, err error) {
	row := repo.db.QueryRow("SELECT * FROM account WHERE username = $1", username)
	err = row.Scan(&a.ID, &a.Name, &a.Username, &a.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			err = model.ErrAccountNotFound
		}
		return
	}

	return
}

func (repo *accountRepo) Update(id string, name string) (err error) {
	q := "UPDATE account SET name = $2 WHERE ID = $1"
	_, err = repo.db.Exec(q, id, name)
	return
}

func (repo *accountRepo) GetAll() (allAccounts []model.Account, err error) {
	query := "SELECT id, name from account"
	rows, err := repo.db.Query(query)
	var accounts []model.Account
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var a model.Account
		err = rows.Scan(&a.ID, &a.Name, &a.Username, &a.Password)
		if err != nil {
			return nil, model.ErrAccountNotFound
		}

		accounts = append(accounts, a)
	}

	return accounts, err
}

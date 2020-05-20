package psql

import (
	"database/sql"

	"sample/model"
)

type sessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) (*sessionRepo, error) {
	query := "CREATE TABLE IF NOT EXISTS session (token varchar primary key, userid varchar, expiry_date timestamp with time zone)"
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &sessionRepo{
		db: db,
	}, nil
}

func (repo *sessionRepo) Add(s model.Session) (err error) {
	q := "INSERT INTO session (token, userid, expiry_date) VALUES ($1, $2, $3)"
	_, err = repo.db.Exec(q, s.Token, s.UserID, s.ExpiryDate)
	return
}

func (repo *sessionRepo) Get(token string) (s model.Session, err error) {
	row := repo.db.QueryRow("SELECT * FROM session WHERE token = $1", token)
	err = row.Scan(&s.Token, &s.UserID, &s.ExpiryDate)
	if err != nil {
		if err == sql.ErrNoRows {
			err = model.ErrSessionNotFound
		}
		return
	}
	return
}

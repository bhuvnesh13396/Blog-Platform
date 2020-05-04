package psql

import (
	"database/sql"
	"sample/model"
)

type tagRepo struct {
	db *sql.DB
}

func NewTagRepo(db *sql.DB) (*tagRepo, error) {
	query := "CREATE TABLE IF NOT EXISTS tag (id varchar primary key, name varchar, description varchar)"
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	query := "CREATE TABLE IF NOT EXISTS tag_article (tag_id varchar NOT NULL, article_id varchar NOT NULL, PRIMARY KEY(tag_id, article_id))"

	return &tagRepo{
		db: db,
	}, nil
}

func (repo *tagRepo) Add(t model.Tag) (err error) {
	query := "INSERT INTO tag (id, name, description) VALUES ($1, $2, $3)"
	_, err = repo.db.Exec(query, t.ID, t.Name, t.Description)
	return
}

package psql

import (
	"database/sql"

	"sample/model"
)

type articleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) (*articleRepo, error) {
	query := "CREATE TABLE IF NOT EXISTS article (id varchar primary key, title varchar, description varchar)"
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &articleRepo{
		db: db,
	}, nil
}

func (repo *articleRepo) Add(a model.Article) (err error) {
	query := "INSERT INTO article (id, title, description) VALUES ($1, $2, $3)"
	_, err = repo.db.Exec(query, a.ID, a.Title, a.Description)
	return
}

func (repo *articleRepo) Get(id string) (a model.Article, err error) {
	row := repo.db.QueryRow("SELECT * FROM article WHERE id = $1", id)
	switch err := row.Scan(&a.ID, &a.Title, &a.Description); err {
	case sql.ErrNoRows:
		return a, model.ErrArticleNotFound
	case nil:
		return a, nil
	}
	return
}

func (repo *articleRepo) Update(id, title string) (err error) {
	query := "UPDATE article SET title = $2 WHERE ID = $1"
	_, err = repo.db.Exec(query, id, title)
	return
}

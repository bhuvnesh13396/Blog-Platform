package psql

import (
	"database/sql"
	"fmt"
	"sample/model"
)

type articleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) (*articleRepo, error) {
	query := "CREATE TABLE IF NOT EXISTS article (id varchar primary key, title varchar, description varchar, user_id varchar)"
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &articleRepo{
		db: db,
	}, nil
}

func (repo *articleRepo) Add(a model.Article) (err error) {
	query := "INSERT INTO article (id, title, description, user_id) VALUES ($1, $2, $3, $4)"
	_, err = repo.db.Exec(query, a.ID, a.Title, a.Description, a.UserID)
	return
}

func (repo *articleRepo) Get(id string) (a model.Article, err error) {
	row := repo.db.QueryRow("SELECT * FROM article WHERE id = $1", id)
	err = row.Scan(&a.ID, &a.Title, &a.Description, &a.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = model.ErrArticleNotFound
		}
		return
	}
	return
}

func (repo *articleRepo) Update(id, title string) (err error) {
	query := "UPDATE article SET title = $2 WHERE ID = $1"
	_, err = repo.db.Exec(query, id, title)
	return
}

func (repo *articleRepo) GetAll() (allArticles []model.Article, err error) {
	query := "SELECT * from article"
	rows, err := repo.db.Query(query)
	var articles []model.Article
	if err != nil {
		return articles, model.ErrArticleNotFound
	}

	defer rows.Close()
	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ID, &a.Title, &a.Description, &a.UserID)
		if err != nil {
			fmt.Printf("%+v\n", a)
			return nil, model.ErrArticleNotFound
		}

		articles = append(articles, a)
	}

	return articles, err
}

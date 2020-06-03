package psql

import (
	"database/sql"
	"sample/model"
)

type tagRepo struct {
	db *sql.DB
}

func NewTagRepo(db *sql.DB) (*tagRepo, error) {
	querytoCreateTag := "CREATE TABLE IF NOT EXISTS tag (id varchar primary key, name varchar)"
	_, err := db.Exec(querytoCreateTag)
	if err != nil {
		return nil, err
	}

	queryToCreateTagArticle := "CREATE TABLE IF NOT EXISTS tag_article (tag_id varchar NOT NULL, article_id varchar NOT NULL, PRIMARY KEY(tag_id, article_id))"
	_, err = db.Exec(queryToCreateTagArticle)

	return &tagRepo{
		db: db,
	}, nil
}

func (repo *tagRepo) Add(t model.Tag) (err error) {
	query := "INSERT INTO tag (id, name) VALUES ($1, $2)"
	_, err = repo.db.Exec(query, t.ID, t.Name)
	return
}

func (repo *tagRepo) AddToArticle(tagID, articleID string) (err error) {
	query := "INSERT INTO tag_article (tag_id, article_id) VALUES ($1, $2)"
	_, err = repo.db.Exec(query, tagID, articleID)
	return
}

func (repo *tagRepo) RemoveFromArticle(tagID, articleID string) (err error) {
	query := "DELETE FROM tag_article WHERE tag_id = $1 AND article_id = $2"
	_, err = repo.db.Exec(query, tagID, articleID)
	return
}

func (repo *tagRepo) GetTagsOnArticle(articleID string) (tags []model.Tag, err error) {
	query := "SELECT id, name from tag where id IN (SELECT tag_id from tag_article where article_id = $1)"
	rows := repo.db.QueryRow(query, articleID)
	var t model.Tag

	switch err := rows.Scan(&t.ID, &t.Name); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		tags = append(tags, t)
	}

	return
}

func (repo *tagRepo) GetArticles(tagID string) (articles []model.Article, err error) {
	query := "SELECT * from article where id IN (SELECT article_id from tag_article WHERE tag_id = $1)"
	rows := repo.db.QueryRow(query, tagID)
	var a model.Article

	switch err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.UserID); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		articles = append(articles, a)
	}

	return
}

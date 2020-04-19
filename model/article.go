package model

import (
	"errors"
)

var (
	ErrArticleNotFound = errors.New("Article not found")
)

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ArticleRepo interface {
	Add(article Article) (err error)
	Get(id string) (article Article, err error)
	Update(id string, title string) (err error)
}

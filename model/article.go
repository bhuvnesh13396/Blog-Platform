package model

import (
	"sample/common/err"
)

var (
	ErrArticleNotFound = err.New(201, "article not found")
)

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UserID      string `json:user_id"`
	Description string `json:"description"`
}

func (a Article) IsValid() bool {
	if len(a.ID) < 1 || len(a.Title) < 1 || len(a.UserID) < 1 || len(a.Description) < 1 {
		return false
	}
	return true
}

type ArticleRepo interface {
	Add(article Article) (err error)
	Get(id string) (article Article, err error)
	Update(id string, title string) (err error)
	GetAll() (article []Article, err error)
}

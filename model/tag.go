package model

import (
	"sample/common/err"
)

var (
	ErrTagNotFound = err.New(501, "Tag not found")
)

type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TagArticle struct {
	TagID     string `json:"tag_id"`
	ArticleID string `json:"article_id"`
}

type TagRepo interface {
	Add(tag Tag) (err error)
	// AddToArticle(tag Tag, articleID string) (err error)
	// RemoveFromArticle(tag Tag, articleID string) (err error)
	// GetTagsOnArticle(articleID string) (tags []Tag, err error)
	// GetArticles(tagID string) (articles []Articles, err error)
}

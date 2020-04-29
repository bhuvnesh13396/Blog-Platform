package model

import (
	"sample/common/err"
	"time"
)

var (
	ErrCommentNotFound = err.New(301, "comment not found")
)

type Comment struct {
	ID          string    `json:"id"`
	Text        string    `json:"text"`
	UserID      string    `json:"user_id"`
	ArticleID   string    `json:"article_id"`
	CreatedDate time.Time `json:"created_date"`
}

type CommentRepo interface {
	Add(comment Comment) (err error)
	Update(id, text string) (err error)
	GetUserComment(userID string) (comments []Comment, err error)
	GetArticleComment(articleID string) (comments []Comment, err error)
}

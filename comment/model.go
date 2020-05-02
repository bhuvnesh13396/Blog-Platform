package comment

import (
	"sample/model"
	"time"
)

type AddReq struct {
	ID          string    `json:"id"`
	Text        string    `json:"text"`
	UserID      string    `json:"user_id"`
	ArticleID   string    `json:"article_id"`
	CreatedDate time.Time `json:"created_date"`
}

type UpdateReq struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type GetRes struct {
	ID          string        `json:"id"`
	Text        string        `json:"text"`
	User        model.Account `json:"account"`
	Article     model.Article `json:"article"`
	CreatedDate time.Time     `json:"created_date"`
}

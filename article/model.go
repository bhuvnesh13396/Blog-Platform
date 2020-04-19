package article

import (
	"sample/model"
)

type AddReq struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UserID      string `json:"user_id"`
	Description string `json:"description"`
}

type UpdateReq struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type GetRes struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	User        model.Account `json:"account"`
}

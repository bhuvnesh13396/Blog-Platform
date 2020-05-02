package article

import "sample/model"

type GetRes struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	User        model.Account `json:"account"`
}

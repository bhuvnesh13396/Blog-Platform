package comment

import (
	"context"
	"sample/common/kit"
	"time"
)

type GetEndpoint kit.Endpoint
type AddEndpoint kit.Endpoint
type UpdateEndpoint kit.Endpoint
type ListEndpoint kit.Endpoint

type Endpoint struct {
	GetEndpoint
	AddEndpoint
	UpdateEndpoint
	ListEndpoint
}

type addRequest struct {
	ID          string    `schema:"id"`
	UserID      string    `schema:"user_id"`
	ArticleID   string    `schema:"article_id"`
	Text        string    `schema:"text"`
	createdDate time.Time `schema:"created_date"`
}

type addResponse struct {
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		err := s.Add(ctx, req.ID, req.UserID, req.ArticleID, req.Text, req.CreatedDate)
		return addResponse{}, err
	}
}

func (e Endpoint) Add(ctx contect.Context, id, userID, articleID, text string, create_date time.Time) {
	request := addRequest{
		ID:          id,
		UserID:      userID,
		ArticleID:   articleID,
		Text:        text,
		CreatedDate: create_date,
	}

	_, err = e(ctx, request)
	return err
}

type updateRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type updateResponse struct {
}

func MakeUpdateEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.Update(ctx, req.ID, req.Text)
		return updateResponse{}, err
	}
}

func (e UpdateEndpoint) Update(ctx context.Context, id, text string) (err error) {
	request := updateRequest{
		ID:   id,
		Text: text,
	}

	_, err = e(ctx, request)
	return err
}

type getUserCommentRequest struct {
	UserID string `schema:"user_id"`
}

type getUserCommentResponse struct {
	Comments []GetRes `json:"comments"`
}

func MakeGetUserCommentEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserCommentRequest)
		comments, err := s.GetUserComment(ctx, req.UserID)
		return getUserCommentResponse{Comments: comments}, err
	}
}

func (e GetEndpoint) GetUserComment(ctx context.Context, user_id string) (comments []GetRes, err error) {
	request := getUserCommentRequest{
		UserID: user_id,
	}
	response, err := e(ctx, request)
	resp := response.(getUserCommentResponse)
	return resp.Comments, err
}

type getArticleCommentRequest struct {
	ArticleID string `schema:"article_id"`
}

type getArticleCommentResponse struct {
	Comments []GetRes `json:"comments"`
}

func MakeGetArticleCommentEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getArticleCommentRequest)
		comments, err := s.GetArticleComment(ctx, req.ArticleID)
		return getArticleCommentResponse{Comments: comments}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, article_id string) (comments []GetRes, err error) {
	request := getArticleCommentRequest{
		ArticleID: article_id,
	}
	response, err := e(ctx, request)
	resp := response.(getArticleCommentResponse)
	return resp.Comments, err
}

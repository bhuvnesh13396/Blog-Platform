package article

import (
	"context"

	"sample/common/kit"
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

type getRequest struct {
	ID string `schema:"id"`
}

type getResponse struct {
	Article GetRes `json:"article"`
}

func MakeGetEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		article, err := s.Get(ctx, req.ID)
		return getResponse{Article: article}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, id string) (article GetRes, err error) {
	request := getRequest{
		ID: id,
	}
	response, err := e(ctx, request)
	resp := response.(getResponse)
	return resp.Article, err
}

type addRequest struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type addResponse struct {
	ID string `json:"id"`
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		id, err := s.Add(ctx, req.UserID, req.Title, req.Description)
		return addResponse{ID: id}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, userID string, title string, description string) (id string, err error) {
	request := addRequest{
		UserID:      userID,
		Title:       title,
		Description: description,
	}
	response, err := e(ctx, request)
	resp := response.(addResponse)
	return resp.ID, err
}

type updateRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type updateResponse struct {
}

func MakeUpdateEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.Update(ctx, req.ID, req.Title)
		return updateResponse{}, err
	}
}

func (e UpdateEndpoint) Update(ctx context.Context, id string, title string) (err error) {
	request := updateRequest{
		ID:    id,
		Title: title,
	}
	_, err = e(ctx, request)
	return err
}

type listRequest struct {
}

type listResponse struct {
	Articles []GetRes `json:"articles"`
}

func MakeListEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		articles, err := s.List(ctx)
		return listResponse{Articles: articles}, err
	}
}

func (e ListEndpoint) List(ctx context.Context) (res []GetRes, err error) {
	request := listRequest{}
	response, err := e(ctx, request)
	resp := response.(listResponse)
	return resp.Articles, err
}

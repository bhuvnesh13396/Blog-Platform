package tag

import (
	"context"
	"sample/common/kit"
	"sample/model"
)

type GetTagsOnArticleEndpoint kit.Endpoint
type GetArticlesEndpoint kit.Endpoint
type AddEndpoint kit.Endpoint
type AddToArticleEndpoint kit.Endpoint
type RemoveFromArticleEndpoint kit.Endpoint

type Endpoint struct {
	GetTagsOnArticleEndpoint
	GetArticlesEndpoint
	AddEndpoint
	AddToArticleEndpoint
	RemoveFromArticleEndpoint
}

type getTagsOnArticleRequest struct {
	ID string `schema:"id"`
}

type getTagsOnArticleResponse struct {
	Tags []model.Tag `json:"tags"`
}

func MakeGetTagsOnArticleEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTagsOnArticleRequest)
		tags, err := s.GetTagsOnArticle(ctx, req.ID)
		return getTagsOnArticleResponse{Tags: tags}, err
	}
}

func (e GetTagsOnArticleEndpoint) GetTagsOnArticle(ctx context.Context, id string) (tags []model.Tag, err error) {
	request := getTagsOnArticleRequest{
		ID: id,
	}
	response, err := e(ctx, request)
	resp := response.(getTagsOnArticleResponse)
	return resp.Tags, err
}

type getArticlesRequest struct {
	ID string `schema:"id"`
}

type getArticleResponse struct {
	Articles []model.Article `json:"articles"`
}

func MakeGetArticlesEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getArticlesRequest)
		articles, err := s.GetArticles(ctx, req.ID)
		return getArticleResponse{Articles: articles}, err
	}
}

func (e GetArticlesEndpoint) GetArticles(ctx context.Context, id string) (Articles []model.Article, err error) {
	request := getArticlesRequest{
		ID: id,
	}
	response, err := e(ctx, request)
	resp := response.(getArticleResponse)
	return resp.Articles, err
}

type addRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type addResponse struct {
	ID string `json:"id"`
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		id, err := s.Add(ctx, req.ID, req.Name)
		return addResponse{ID: id}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, id, name string) (interface{}, error) {
	request := addRequest{
		ID:   id,
		Name: name,
	}
	response, err := e(ctx, request)
	resp := response.(addResponse)
	return resp.ID, err
}

type addToArticleRequest struct {
	TagID     string `json:"tag_id"`
	ArticleID string `json:"article_id"`
}

type addToArticleResponse struct {
	ID string `json:"id"`
}

func MakeAddToArticleEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addToArticleRequest)
		id, err := s.AddToArticle(ctx, req.TagID, req.ArticleID)
		return addToArticleResponse{ID: id}, err
	}
}

func (e AddToArticleEndpoint) AddToArticle(ctx context.Context, tagID, articleID string) (id string, err error) {
	request := addToArticleRequest{
		TagID:     tagID,
		ArticleID: articleID,
	}

	response, err := e(ctx, request)
	resp := response.(addToArticleResponse)
	return resp.ID, err

}

type removeFromArticleRequest struct {
	ID string `json:"id"`
}

type removeFromArticleResponse struct {
}

func MakeRemoveFromArticleEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(removeFromArticleRequest)
		err := s.RemoveFromArticle(ctx, req.ID)
		return removeFromArticleResponse{}, err
	}
}

func (e RemoveFromArticleEndpoint) RemoveFromArticle(ctx context.Context, id string) (err error) {
	request := removeFromArticleRequest{
		ID: id,
	}

	_, err = e(ctx, request)
	return err
}

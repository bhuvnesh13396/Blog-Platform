package client

import (
	"net/http"
	"net/url"
	"sample/article"
	"sample/common/auth/token"
	"sample/common/kit"
	"sample/tag"
)

func New(instance string, client *http.Client) (tag.Service, error){
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	opts := []kit.ClientOption{
		kit.SetClient(client),
		kit.ClientBefore(token.HTTPTokenFromContext)
	}

	addEndpoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/tag/add")
		kit.EncodeJSONRequest,
		tag.DecodeAddResponse,
		opts...,
	).Endpoint()

	addToArticleEndpoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/tag/addToArticle"),
		kit.EncodeJSONRequest,
		tag.DecodeAddToArticleResponse,
		opts...,
	).Endpoint()

	getTagsOnArticleEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/tag/onArticle"),
		kit.EncodeSchemaRequest,
		tag.DecodeGetTagsOnArticleResponse,
		opts...,
	).Endpoint()

	getArticlesEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/tag/allArticles"),
		kit.EncodeSchemaRequest,
		tag.DecodeGetArticleResponse,
		opts...,
	).Endpoint()

	removeFromArticleEndpoint := kit.NewClient(
		http.MethodDelete,
		copyURL(u, "/tag/removeFromArticle"),
		kit.EncodeJSONRequest,
		tag.DecodeDeleteFromArticleResponse,
		opts...,
	).Endpoint()

	return &tag.Endpoint{
		GetTagsOnArticleEndpoint: tag.GetTagsOnArticleEndpoint(getTagsOnArticleEndpoint),
		GetArticlesEndpoint: tag.GetArticlesEndpoint(getArticlesEndpoint),
		AddEndpoint: tag.AddEndpoint(addEndpoint),
		AddToArticleEndpoint: tag.AddToArticleEndpoint(addToArticleEndpoint),
		RemoveFromArticleEndpoint: tag.RemoveFromArticleEndpoint(removeFromArticleEndpoint),
	}, nil
}

func copyURL(u *url.URL, path string) *url.URL {
	c := *u
	c.Path = path
	return &c
}
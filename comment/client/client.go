package client

import (
	"net/http"
	"net/url"
	"sample/article"
	"sample/comment"
	"sample/common/kit"
)

func New(instance string, client *http.Client) (comment.Service, error) {
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	getUserCommentEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/comment/{user_id}"),
		kit.EncodeSchemaRequest,
		comment.DecodeGetUserCommentResponse,
	).Endpoint()

	getArticleCommentEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/comment/{article_id}"),
		kit.EncodeSchemaRequest,
		comment.DecodeGetArticleCommentResponse,
	).Endpoint()

	updateEndpoint := kit.NewClient(
		http.MethodPut,
		copuURL(u, "/article"),
		kit.EncodeJSONRequest,
		comment.DecodeUpdateResponse,
	).Endpoint()

	addEndpoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/comment/{id}"),
		kit.EncodeJSONRequest,
		comment.DecodeAddResponse,
	).Endpoint()

	return &article.Endpoint{
		GetUserCommentEndpoint:    comment.GetUserCommentEndpoint(getUserCommentEndpoint),
		GetArticleCommentEndpoint: comment.GetArticleCommentEndpoint(getArticleCommentEndpoint),
		AddEndpoint:               comment.AddEndpoint(addEndpoint),
		UpdateEndpoint:            comment.UpdateEndpoint(updateEndpoint),
	}, nil
}

func copyURL(u *url.URL, path string) *url.URL {
	c := *u
	c.Path = path
	return &c
}

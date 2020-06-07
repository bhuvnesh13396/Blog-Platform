package tag

import (
	"context"
	"encoding/json"
	"net/http"
	"sample/common/auth/token"
	"sample/common/kit"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	ctx = context.Background()
)

func NewHandler(s Service) http.Handler {
	r := mux.NewRouter()

	opts := []kit.ServerOption{
		kit.ServerBefore(token.HTTPTokenFromContext),
	}

	add := kit.NewServer(
		MakeAddEndpoint(s),
		DecodeAddRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	addToArticle := kit.NewServer(
		MakeAddToArticleEndpoint(s),
		DecodeAddToArticleRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	removeFromArticle := kit.NewServer(
		MakeRemoveFromArticleEndpoint(s),
		DecodeDeleteFromArticleRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	getTagsOnArticle := kit.NewServer(
		MakeGetTagsOnArticleEndpoint(s),
		DecodeGetTagsOnArticleRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	getArticle := kit.NewServer(
		MakeGetArticlesEndpoint(s),
		DecodeGetArticleRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	r.Handle("/tag/allArticles", getArticle).Methods(http.MethodGet)
	r.Handle("/tag/onArticle", getTagsOnArticle).Methods(http.MethodGet)
	r.Handle("/tag/removeFromArticle", removeFromArticle).Methods(http.MethodDelete)
	r.Handle("/tag", add).Methods(http.MethodPut)
	r.Handle("/tag/addToArticle", addToArticle).Methods(http.MethodPut)

	return r
}

func DecodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeAddResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp addResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeAddToArticleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addToArticleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeAddToArticleResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp addToArticleResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeDeleteFromArticleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req removeFromArticleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeDeleteFromArticleResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp removeFromArticleResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeGetTagsOnArticleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getTagsOnArticleRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeGetTagsOnArticleResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp getTagsOnArticleResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeGetArticleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getArticlesRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeGetArticleResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp getArticleResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

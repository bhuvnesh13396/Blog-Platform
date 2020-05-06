package comment

import (
	"context"
	"encoding/json"
	"net/http"
	"sample/common/kit"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	ctx = context.Background()
)

func NewHandler(s Service) http.Handler {
	r := mux.NewRouter()

	getUserComment := kit.NewServer(
		MakeGetUserCommentEndpoint(s),
		DecodeGetUserCommentRequest,
		kit.EncodeJSONResponse,
	)

	getArticleComment := kit.NewServer(
		MakeGetArticleCommentEndpoint(s),
		DecodeGetArticleCommentRequest,
		kit.EncodeJSONResponse,
	)

	add := kit.NewServer(
		MakeAddEndpoint(s),
		DecodeAddRequest,
		kit.EncodeJSONResponse,
	)

	update := kit.NewServer(
		MakeUpdateEndpoint(s),
		DecodeUpdateRequest,
		kit.EncodeJSONResponse,
	)

	r.Handle("/comment/{user_id}", getUserComment).Methods(http.MethodGet)
	r.Handle("/comment/{article_id}", getArticleComment).Methods(http.MethodGet)
	r.Handle("/comment", add).Methods(http.MethodPost)
	r.Handle("/comment", update).Methods(http.MethodPut)

	return r

}

func DecodeGetUserCommentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeGetUserCommentResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp getResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeGetArticleCommentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeGetArticleCommentResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp getResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
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

func DecodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req updateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeUpdateResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp updateResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

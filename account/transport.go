package account

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

	get := kit.NewServer(
		MakeGetEndpoint(s),
		DecodeGetRequest,
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

	list := kit.NewServer(
		MakeListEndpoint(s),
		DecodeListRequest,
		kit.EncodeJSONResponse,
	)

	r.Handle("/account/{id}", get).Methods(http.MethodGet)
	r.Handle("/account", add).Methods(http.MethodPost)
	r.Handle("/account", update).Methods(http.MethodPut)
	r.Handle("/account", list).Methods(http.MethodGet)

	return r
}

func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeAddResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var req addRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
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

func DecodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req listRequest
	return req, nil
}

func DecodeListResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp listResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

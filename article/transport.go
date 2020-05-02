package article

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

	r.Handle("/article/one", get).Methods(http.MethodGet)
	r.Handle("/article", add).Methods(http.MethodPost)
	r.Handle("/article", update).Methods(http.MethodPut)
	r.Handle("/article", list).Methods(http.MethodGet)

	return r
}

func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeGetResponse(ctx context.Context, r *http.Response) (interface{}, error) {
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

func DecodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req listRequest
	return req, nil
}

func DecodeListResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp listResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

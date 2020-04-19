package account

import (
	"context"
	"encoding/json"
	"net/http"

	"sample/common/resp"

	"github.com/gorilla/mux"
)

var (
	ctx = context.Background()
)

func NewHandler(s Service) http.Handler {
	r := mux.NewRouter()

	getAccount := func(w http.ResponseWriter, req *http.Request) {
		id, _ := mux.Vars(req)["id"]
		a, err := s.GetAccount(ctx, id)
		resp.WriteResp(w, a, err)
	}

	addAccount := func(w http.ResponseWriter, req *http.Request) {
		var addReq AddReq
		err := json.NewDecoder(req.Body).Decode(&addReq)
		if err != nil {
			resp.WriteResp(w, nil, err)
			return
		}
		err = s.AddAccount(ctx, addReq.ID, addReq.Name)
		resp.WriteResp(w, nil, err)
	}

	updateAccount := func(w http.ResponseWriter, req *http.Request) {
		var updateReq UpdateReq
		err := json.NewDecoder(req.Body).Decode(&updateReq)
		if err != nil {
			resp.WriteResp(w, nil, err)
			return
		}

		err = s.UpdateAccount(ctx, updateReq.ID, updateReq.Name)
		resp.WriteResp(w, nil, err)
	}

	r.HandleFunc("/account/{id}", getAccount).Methods(http.MethodGet)
	r.HandleFunc("/account", addAccount).Methods(http.MethodPost)
	r.HandleFunc("/account", updateAccount).Methods(http.MethodPut)

	return r
}

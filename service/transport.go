package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ctx = context.Background()
)

type resp struct {
	Data  *interface{} `json:"data"`
	Error string       `json:"error,omitempty"`
}

func writeResp(w http.ResponseWriter, data interface{}, err error) {
	var r resp
	if err == nil {
		r.Data = &data
	} else {
		r.Error = err.Error()
	}
	json.NewEncoder(w).Encode(r)
}

func NewHandler(s Service) http.Handler {
	r := mux.NewRouter()

	getAccount := func(w http.ResponseWriter, req *http.Request) {
		id, _ := mux.Vars(req)["id"]
		a, err := s.GetAccount(ctx, id)
		writeResp(w, a, err)
	}

	addAccount := func(w http.ResponseWriter, req *http.Request) {
		var addReq AddReq
		err := json.NewDecoder(req.Body).Decode(&addReq)
		if err != nil {
			writeResp(w, nil, err)
			return
		}
		err = s.AddAccount(ctx, addReq.ID, addReq.Name)
		writeResp(w, nil, err)
	}

	updateAccount := func(w http.ResponseWriter, req *http.Request) {
		var addReq AddReq
		err := json.NewDecoder(req.Body).Decode(&addReq)
		if err != nil {
			writeResp(w, nil, err)
			return
		}

		err = s.UpdateAccount(ctx, addReq.ID, addReq.Name)
		writeResp(w, nil, err)
	}

	r.HandleFunc("/account/{id}", getAccount).Methods(http.MethodGet)
	r.HandleFunc("/account", addAccount).Methods(http.MethodPost)
	r.HandleFunc("/account/{id}", updateAccount).Methods(http.MethodPut)

	return r
}

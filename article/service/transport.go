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
	getArticle := func(w http.ResponseWriter, req *http.Request) {
		id, _ := mux.Vars(req)["id"]
		a, err := s.GetArticle(ctx, id)
		writeResp(w, a, err)
	}

	addArticle := func(w http.ResponseWriter, req *http.Request) {
		var addReq AddReq
		err := json.NewDecoder(req.Body).Decode(&addReq)
		if err != nil {
			writeResp(w, nil, err)
			return
		}
		err = s.AddArticle(ctx, addReq.ID, addReq.Title, addReq.Description)
		writeResp(w, nil, err)
	}

	updateArticle := func(w http.ResponseWriter, req *http.Request) {
		var updateReq UpdateReq
		err := json.NewDecoder(req.Body).Decode(&updateReq)
		if err != nil {
			writeResp(w, nil, err)
			return
		}

		err = s.UpdateArticle(ctx, updateReq.ID, updateReq.Title)
		writeResp(w, nil, err)
	}

	r.HandleFunc("/article", addArticle).Methods(http.MethodPost)
	r.HandleFunc("/article", updateArticle).Methods(http.MethodPut)
	r.HandleFunc("/article/{id}", getArticle).Methods(http.MethodGet)
	return r
}

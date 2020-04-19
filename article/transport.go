package article

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

	getArticle := func(w http.ResponseWriter, req *http.Request) {
		id, _ := mux.Vars(req)["id"]
		a, err := s.GetArticle(ctx, id)
		resp.WriteResp(w, a, err)
	}

	addArticle := func(w http.ResponseWriter, req *http.Request) {
		var addReq AddReq
		err := json.NewDecoder(req.Body).Decode(&addReq)
		if err != nil {
			resp.WriteResp(w, nil, err)
			return
		}
		err = s.AddArticle(ctx, addReq.ID, addReq.UserID, addReq.Title, addReq.Description)
		resp.WriteResp(w, nil, err)
	}

	updateArticle := func(w http.ResponseWriter, req *http.Request) {
		var updateReq UpdateReq
		err := json.NewDecoder(req.Body).Decode(&updateReq)
		if err != nil {
			resp.WriteResp(w, nil, err)
			return
		}

		err = s.UpdateArticle(ctx, updateReq.ID, updateReq.Title)
		resp.WriteResp(w, nil, err)
	}

	r.HandleFunc("/article", addArticle).Methods(http.MethodPost)
	r.HandleFunc("/article", updateArticle).Methods(http.MethodPut)
	r.HandleFunc("/article/{id}", getArticle).Methods(http.MethodGet)
	return r
}

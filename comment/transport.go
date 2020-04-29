package comment

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

	addComment := func(w http.ResponseWriter, req *http.Request) {
		var addReq AddReq
		err := json.NewDecoder(req.Body).Decode(&addReq)
		if err != nil {
			resp.WriteResp(w, nil, err)
			return
		}

		err = s.Add(ctx, addReq.ID, addReq.UserID, addReq.Text, addReq.ArticleID, addReq.CreatedDate)
		resp.WriteResp(w, nil, err)
	}

	updateComment := func(w http.ResponseWriter, req *http.Request) {
		var updateReq UpdateReq
		err := json.NewDecoder(req.Body).Decode(&updateReq)
		if err != nil {
			resp.WriteResp(w, nil, err)
			return
		}

		err = s.Update(ctx, updateReq.ID, updateReq.Text)
		resp.WriteResp(w, nil, err)
	}

	getUserComment := func(w http.ResponseWriter, req *http.Request) {
		id, _ := mux.Vars(req)["id"]
		c, err := s.GetUserComment(ctx, id)
		resp.WriteResp(w, c, err)
	}

	getArticleComment := func(w http.ResponseWriter, req *http.Request) {
		id, _ := mux.Vars(req)["id"]
		c, err := s.GetArticleComment(ctx, id)
		resp.WriteResp(w, c, err)
	}

	r.HandleFunc("/comment", addComment).Methods(http.MethodPost)
	r.HandleFunc("/comment", updateComment).Methods(http.MethodPut)
	r.HandleFunc("/comment/user/id", getUserComment).Methods(http.MethodGet)
	r.HandleFunc("/comment/article/id", getArticleComment).Methods(http.MethodGet)

	return r
}

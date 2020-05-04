package tag

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

	addTag := func(w http.ResponseWriter, req *http.Request) {
		var addReq AddReq
		err := json.NewDecoder(req.Body).Decode(&addReq)
		if err != nil {
			resp.WriteResp(w, nil, err)
			return
		}
		err = s.Add(ctx, addReq.ID, addReq.Name, addReq.Descritption)
		resp.WriteResp(w, nil, err)
	}

	r.HandleFunc("/tag", addTag).Methods(http.MethodPost)
	return r
}

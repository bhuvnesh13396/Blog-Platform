package kit

import (
	"net/http"
)

type server struct {
	e   Endpoint
	dec DecodeRequestFunc
	enc EncodeResponseFunc
}

func NewServer(e Endpoint, dec DecodeRequestFunc, enc EncodeResponseFunc) *server {
	return &server{
		e:   e,
		dec: dec,
		enc: enc,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := s.dec(ctx, r)
	if err != nil {
		s.enc(ctx, w, nil, err)
		return
	}

	resp, err := s.e(ctx, req)

	err = s.enc(ctx, w, resp, err)
	if err != nil {
		return
	}
}

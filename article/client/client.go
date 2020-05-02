package client

import (
	"net/http"
	"net/url"

	"sample/article"
	"sample/common/kit"
)

func New(instance string, client *http.Client) (article.Service, error) {
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	getEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/article/one"),
		kit.EncodeSchemaRequest,
		article.DecodeGetResponse,
	).Endpoint()

	addEndpoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/article"),
		kit.EncodeJSONRequest,
		article.DecodeAddResponse,
	).Endpoint()

	updateEndpoint := kit.NewClient(
		http.MethodPut,
		copyURL(u, "/article"),
		kit.EncodeJSONRequest,
		article.DecodeUpdateResponse,
	).Endpoint()

	listEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/article"),
		kit.EncodeSchemaRequest,
		article.DecodeListResponse,
	).Endpoint()

	return &article.Endpoint{
		GetEndpoint:    article.GetEndpoint(getEndpoint),
		AddEndpoint:    article.AddEndpoint(addEndpoint),
		UpdateEndpoint: article.UpdateEndpoint(updateEndpoint),
		ListEndpoint:   article.ListEndpoint(listEndpoint),
	}, nil
}

func copyURL(u *url.URL, path string) *url.URL {
	c := *u
	c.Path = path
	return &c
}

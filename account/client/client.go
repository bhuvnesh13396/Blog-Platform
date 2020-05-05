package client

import(
	"net/http"
	"net/url"
	"sample/article"
	"sample/common/kit"
)

func New(instance string, client *http.Client) (account.Service, error){
	u, err := url.Parse(instance)
	if err != nil{
		return nil, err
	}

	getEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/account/{id}"),
		kit.EncodeSchemaRequest,
		account.DecodeGetResponse,
	).Endpoint()

	addEndPoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/account")
		kit.EncodeJSONRequest,
		account.DecodeAddResponse,
	).Endpoint()

	updateEndpoint := kit.NewClient(
		http.MethodPut,
		copyURL(u, "/account")
		kit.EncodeJSONRequest,
		account.DecodeUpdateResponse,
	).Endpoint()

	listEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/account"),
		kit.EncodeSchemaRequest,
		account.DecodeListResponse,
	).Endpoint()

	return &account.Endpoint{
		GetEndpoint	:	account.GetEndpoint(getEndpoint),
		AddEndpoint	:	account.AddEndpoint(addEndPoint),
		UpdateEndpoint	:	account.UpdateEndpoint(updateEndpoint),
		listEndpoint	:	account.ListEndpoint(listEndpoint),
	}, nil
}

func copyURL(u *url.URL, path string) *url.URL{
	c := *u
	c.Path = path
	return &c
}
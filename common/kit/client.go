package kit

import (
	"context"
	"net/http"
	"net/url"
)

type client struct {
	client *http.Client
	method string
	url    *url.URL
	enc    EncodeRequestFunc
	dec    DecodeResponseFunc
}

func NewClient(method string, url *url.URL, enc EncodeRequestFunc, dec DecodeResponseFunc) *client {
	return &client{
		client: http.DefaultClient,
		method: method,
		url:    url,
		enc:    enc,
		dec:    dec,
	}
}

func (c *client) Endpoint() Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		hreq, err := http.NewRequest(c.method, c.url.String(), nil)
		if err != nil {
			return nil, err
		}

		err = c.enc(ctx, hreq, request)
		if err != nil {
			return nil, err
		}

		hres, err := c.client.Do(hreq)
		if err != nil {
			return nil, err
		}

		defer hres.Body.Close()

		return c.dec(ctx, hres)
	}
}

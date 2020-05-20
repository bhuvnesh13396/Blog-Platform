package client

import (
	"net/http"
	"net/url"

	"sample/auth"
	"sample/common/auth/token"
	"sample/common/kit"
)

func New(instance string, client *http.Client) (auth.Service, error) {
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	opts := []kit.ClientOption{
		kit.SetClient(client),
		kit.ClientBefore(token.HTTPTokenFromContext),
	}

	signinEndpoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/auth/signin"),
		kit.EncodeJSONRequest,
		auth.DecodeSigninResponse,
		opts...,
	).Endpoint()

	verifyTokenEndpoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/auth/verify"),
		kit.EncodeJSONRequest,
		auth.DecodeVerifyTokenResponse,
		opts...,
	).Endpoint()

	return &auth.Endpoint{
		SigninEndpoint:      auth.SigninEndpoint(signinEndpoint),
		VerifyTokenEndpoint: auth.VerifyTokenEndpoint(verifyTokenEndpoint),
	}, nil
}

func copyURL(u *url.URL, path string) *url.URL {
	c := *u
	c.Path = path
	return &c
}

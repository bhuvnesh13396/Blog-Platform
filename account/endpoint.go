package account

import (
	"context"
	"sample/common/kit"
	"sample/model"
)

type GetEndpoint kit.Endpoint
type AddEndpoint kit.Endpoint
type UpdateEndpoint kit.Endpoint
type ListEndpoint kit.Endpoint

type Endpoint struct {
	GetEndpoint
	AddEndpoint
	UpdateEndpoint
	ListEndpoint
}

type getRequest struct {
	ID string `schema:"id"`
}

type getResponse struct {
	Account model.Account `json:"account"`
}

func MakeGetEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		account, err := s.GetAccount(ctx, req.ID)
		return getResponse{Account: account}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, id string) (account model.Account, err error) {
	request := getRequest{
		ID: id,
	}
	response, err := e(ctx, request)
	resp := response.(getResponse)
	return resp.Account, err
}

type addRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type addResponse struct {
}

func (e AddEndpoint) Add(ctx context.Context, id, name string) (err error) {
	request := addRequest{
		ID:   id,
		Name: name,
	}
	_, err = e(ctx, request)
	return err
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		err := s.AddAccount(ctx, req.ID, req.Name)
		return addResponse{}, err
	}
}

type updateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type updateResponse struct {
}

func MakeUpdateEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.UpdateAccount(ctx, req.ID, req.Name)
		return updateResponse{}, err
	}
}

func (e UpdateEndpoint) Update(ctx context.Context, id, name string) (err error) {
	request := updateRequest{
		ID:   id,
		Name: name,
	}

	_, err = e(ctx, request)
	return err
}

type listRequest struct {
}

type listResponse struct {
	Accounts []model.Account `json:"accounts"`
}

func MakeListEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		accounts, err := s.GetAllAccount(ctx)
		return listResponse{Accounts: accounts}, err
	}
}

func (e ListEndpoint) List(ctx context.Context) (res []model.Account, err error) {
	request := getRequest{}
	response, err := e(ctx, request)
	resp := response.(listResponse)
	return resp.Accounts, err
}

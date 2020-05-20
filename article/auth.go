package article

import (
	"context"

	"sample/auth"
	"sample/common/auth/token"
	"sample/common/err"
)

var (
	errAccessTokenNotFound = err.New(1, "access token not found")
)

type authSvc struct {
	Service
	authService auth.Service
}

func NewAuthService(s Service, authService auth.Service) Service {
	return &authSvc{
		Service:     s,
		authService: authService,
	}
}

func (s *authSvc) verifyToken(ctx context.Context) (userID string, err error) {
	token, ok := ctx.Value(token.ContextKey).(string)
	if !ok || len(token) < 1 {
		return "", errAccessTokenNotFound
	}

	return s.authService.VerifyToken(ctx, token)
}

func (s *authSvc) Get(ctx context.Context, id string) (article GetRes, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Get(ctx, id)
}

func (s *authSvc) Add(ctx context.Context, userID string, title string, description string) (id string, err error) {
	userID, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Add(ctx, userID, title, description)
}

func (s *authSvc) Update(ctx context.Context, id string, title string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Update(ctx, id, title)
}

func (s *authSvc) List(ctx context.Context) (res []GetRes, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.List(ctx)
}

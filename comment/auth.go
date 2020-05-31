package comment

import (
	"context"
	"sample/auth"
	"sample/common/auth/token"
	"sample/common/err"
	"time"
)

var (
	errAccessTokenNotFound = err.New(3, "Access token not found.")
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

func (s *authSvc) GetUserComment(ctx context.Context, username string) (comment []GetRes, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}
	return s.Service.GetUserComment(ctx, username)
}

func (s *authSvc) Add(ctx context.Context, id, userID, articleID, text string, createdDate time.Time) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}
	return s.Service.Add(ctx, id, userID, articleID, text, createdDate)
}

func (s *authSvc) GetArticleComment(ctx context.Context, articleID string) (comments []GetRes, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.GetArticleComment(ctx, articleID)
}

func (s *authSvc) Update(ctx context.Context, id, text string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Update(ctx, id, text)
}

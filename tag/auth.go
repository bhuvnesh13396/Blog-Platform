package tag

import (
	"context"
	"sample/auth"
	"sample/common/auth/token"
	"sample/common/err"
	"sample/model"
)

var (
	errAccessTokenNotFound = err.New(4, "Access token not found")
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

func (s *authSvc) Add(ctx context.Context, id, name string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}
	return s.Service.Add(ctx, id, name)
}

func (s *authSvc) AddToArticle(ctx context.Context, tagID, articleID string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}
	return s.Service.AddToArticle(ctx, tagID, articleID)
}

func (s *authSvc) RemoveFromArticle(ctx context.Context, tagID, articleID string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}
	return s.Service.RemoveFromArticle(ctx, tagID, articleID)
}

func (s *authSvc) GetTagsOnArticle(ctx context.Context, articleID string) (tags []model.Tag, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}
	return s.Service.GetTagsOnArticle(ctx, articleID)

}

func (s *authSvc) GetArticles(ctx context.Context, tagID string) (articles []model.Article, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}
	return s.Service.GetArticles(ctx, tagID)
}

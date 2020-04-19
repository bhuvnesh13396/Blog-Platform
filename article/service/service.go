package service

import (
	"context"
	"errors"
	"sample/article/model"
)

var (
	errInvalidArgument = errors.New("invalid argument")
	errInvalidId       = errors.New("No such article exists")
)

type Service interface {
	GetArticle(ctx context.Context, id string) (article model.Article, err error)
	AddArticle(ctx context.Context, id string, title string, description string) (err error)
	UpdateArticle(ctx context.Context, id string, title string) (err error)
	//GetAllArticle(ctx context.Context) (err error)
}

type service struct {
	articleRepo model.ArticleRepo
}

func NewService(articleRepo model.ArticleRepo) Service {
	return &service{
		articleRepo: articleRepo,
	}
}

func (s *service) GetArticle(ctx context.Context, id string) (article model.Article, err error) {
	if len(id) < 1 {
		err = errInvalidArgument
		return
	}
	return s.articleRepo.Get(id)
}

func (s *service) AddArticle(ctx context.Context, id string, title string, description string) (err error) {
	if len(id) < 1 || len(title) < 1 || len(description) < 1 {
		err = errInvalidArgument
		return
	}

	article := model.Article{
		ID:          id,
		Title:       title,
		Description: description,
	}

	return s.articleRepo.Add(article)
}

func (s *service) UpdateArticle(ctx context.Context, id string, title string) (err error) {
	if len(id) < 1 || len(title) < 1 {
		err = errInvalidArgument
		return
	}

	return s.articleRepo.Update(id, title)
}

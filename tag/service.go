package tag

import (
	"context"
	"sample/common/err"
	"sample/model"
)

var (
	errInvalidArgument = err.New(502, "Invalid argument")
)

type Service interface {
	Add(ctx context.Context, id, name, description string) (err error)
	// AddToArticle(ctx context.Context, tagID, articleID string) (err error)
	// RemoveFromArticle(ctx context.Context, tagID, articleID string) (err error)
	// GetTagsOnArticle(ctx context.Context, articleID string) (tags []model.Tag, err error)
	// GetArticles(ctx context.Context, tagID string) (articles []model.Articles, err error)
}

type service struct {
	tagRepo model.TagRepo
}

func NewService(tagRepo model.TagRepo) Service {
	return &service{
		tagRepo: tagRepo,
	}
}

func (s *service) Add(ctx context.Context, id, name, description string) (err error) {
	if len(id) < 1 || len(name) < 1 || len(description) < 1 {
		err = errInvalidArgument
		return
	}

	t := model.Tag{
		ID:          id,
		Name:        name,
		Description: description,
	}

	return s.tagRepo.Add(t)
}

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
	AddToArticle(ctx context.Context, tagID, articleID string) (err error)
	RemoveFromArticle(ctx context.Context, tagID, articleID string) (err error)
	GetTagsOnArticle(ctx context.Context, articleID string) (tags []model.Tag, err error)
	GetArticles(ctx context.Context, tagID string) (articles []model.Article, err error)
}

type service struct {
	tagRepo     model.TagRepo
	articleRepo model.ArticleRepo
}

func NewService(tagRepo model.TagRepo, articleRepo model.ArticleRepo) Service {
	return &service{
		tagRepo:     tagRepo,
		articleRepo: articleRepo,
	}
}

func (s *service) Add(ctx context.Context, id, name, description string) (err error) {
	if len(id) < 1 || len(name) < 1 || len(description) < 1 {
		err = errInvalidArgument
		return
	}

	t := model.Tag{
		ID:   id,
		Name: name,
	}

	return s.tagRepo.Add(t)
}

func (s *service) AddToArticle(ctx context.Context, id string, name string, tagID string, articleID string) (err error) {
	if len(tagID) < 1 || len(articleID) < 1 {
		return errInvalidArgument
	}

	tag := model.Tag{
		ID:   id,
		Name: name,
	}

	return s.tagRepo.AddToArticle(tag)
}

func (s *service) RemoveFromArticle(ctx context.Context, tagID string, articleID string) (err error) {
	if len(tagID) < 1 || len(articleID) < 1 {
		err = errInvalidArgument
		return
	}

	return s.commentRepo.RemoveFromArticle(tagID, articleID)
}

func (s *service) GetTagsOnArticle(ctx context.Context, articleID string) (tags []model.Tag, err error) {
	if len(articleID) < 1 {
		err = errInvalidArgument
		return nil, err
	}

	tagsOnArticle, err := s.tagRepo.GetTagsOnArticle(articleID)
	if err != nil {
		return nil, err
	}

	for i := range tagsOnArticle {
		tag, err := &tagsOnArticle[i]

		article, err := s.articleRepo.Get(tag.ArticleID)
		if err != nil {
			return nil, err
		}

		t := model.Tag{
			ID:      tag.Id,
			Article: article,
		}

		tags = append(tags, t)
	}

	return
}

func (s *service) GetArticles(ctx context.Context, tagID string) (articles []model.Article, err error) {
	if len(tagID) < 1 {
		err = errInvalidArgument
		return nil, err
	}

	articles, err = s.articleRepo.Get(tagID)
	if err != nil {
		return nil, err
	}

	return
}

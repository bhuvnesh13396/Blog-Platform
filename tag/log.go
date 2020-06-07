package tag

import (
	"context"
	"sample/common/kit"
	"sample/model"
	"time"
)

type logSvc struct {
	Service
	logger kit.Logger
}

func NewLogService(s Service, logger kit.Logger) Service {
	return &logSvc{
		Service: s,
		logger:  logger,
	}
}

func (s *logSvc) Add(ctx context.Context, id, name string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Add",
			"id", id,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())

	return s.Service.Add(ctx, id, name)
}

func (s *logSvc) AddToArticle(ctx context.Context, tagID, articleID string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "AddToArticle",
			"tagId", tagID,
			"articleID", articleID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())

	return s.Service.AddToArticle(ctx, tagID, articleID)
}

func (s *logSvc) RemoveFromArticle(ctx context.Context, tagID, articleID string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "RemoveFromArticle",
			"tagID", tagID,
			"articleID", articleID,
			"took", time.Since(t),
			"err", err,
		)

	}(time.Now())

	return s.Service.RemoveFromArticle(ctx, tagID, articleID)
}

func (s *logSvc) GetTagsOnArticle(ctx context.Context, articleID string) (tags []model.Tag, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "GetTagsOnArticle",
			"articleID", articleID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())

	return s.Service.GetTagsOnArticle(ctx, articleID)

}

func (s *logSvc) GetArticles(ctx context.Context, tagID string) (articles []model.Article, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "GetArticles",
			"tagID", tagID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())

	return s.Service.GetArticles(ctx, tagID)
}

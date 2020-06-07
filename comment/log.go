package comment

import (
	"context"
	"sample/common/kit"
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

func (s *logSvc) Add(ctx context.Context, id, userID, articleID, text string, createdDate time.Time) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Add",
			"id", id,
			"userID", userID,
			"articleID", articleID,
			"text", text,
			"createdDate", createdDate,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Add(ctx, id, userID, articleID, text, createdDate)
}

func (s *logSvc) GetUserComment(ctx context.Context, userID string) (comments []GetRes, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "GetUserComment",
			"userID", userID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetUserComment(ctx, userID)
}

func (s *logSvc) GetArticleComment(ctx context.Context, articleID string) (comments []GetRes, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "GetArticleComment",
			"articleID", articleID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetArticleComment(ctx, articleID)
}

func (s *logSvc) Update(ctx context.Context, id, text string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Update",
			"id", id,
			"text", text,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Update(ctx, id, text)
}

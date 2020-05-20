package article

import (
	"context"
	"time"

	"sample/common/kit"
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

func (s *logSvc) Get(ctx context.Context, id string) (article GetRes, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Get",
			"id", id,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Get(ctx, id)
}

func (s *logSvc) Add(ctx context.Context, userID string, title string, description string) (id string, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Add",
			"user_id", userID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Add(ctx, userID, title, description)
}

func (s *logSvc) Update(ctx context.Context, id string, title string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Update",
			"id", "id",
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Update(ctx, id, title)
}

func (s *logSvc) List(ctx context.Context) (res []GetRes, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "List",
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.List(ctx)
}

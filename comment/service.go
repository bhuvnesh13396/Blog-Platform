package comment

import (
	"context"
	"sample/common/err"
	"sample/model"
	"time"
)

var (
	errInvalidArgument = err.New(503, "invalid argument")
)

type Service interface {
	Add(ctx context.Context, id, userID, articleID, text string, createdDate time.Time) (err error)
	GetUserComment(ctx context.Context, userID string) (comments []GetRes, err error)
	GetArticleComment(ctx context.Context, articleID string) (comments []GetRes, err error)
	Update(ctx context.Context, id, text string) (err error)
}

type service struct {
	commentRepo model.CommentRepo
	articleRepo model.ArticleRepo
	accountRepo model.AccountRepo
}

func NewService(commentRepo model.CommentRepo, articleRepo model.ArticleRepo, accountRepo model.AccountRepo) Service {
	return &service{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
		accountRepo: accountRepo,
	}
}

func (s *service) Add(ctx context.Context, id, userID, articleID, text string, createdDate time.Time) (err error) {
	if len(id) < 1 || len(userID) < 1 || len(articleID) < 1 || len(text) < 1 {
		return errInvalidArgument
	}

	comment := model.Comment{
		ID:          id,
		Text:        text,
		UserID:      userID,
		ArticleID:   articleID,
		CreatedDate: createdDate,
	}

	return s.commentRepo.Add(comment)
}

func (s *service) GetUserComment(ctx context.Context, userId string) (comments []GetRes, err error) {
	if len(userId) < 1 {
		err = errInvalidArgument
		return
	}

	allUserComments, err := s.commentRepo.GetUserComment(userId)
	if err != nil {
		return
	}

	for i := range allUserComments {
		c := &allUserComments[i]
		user, err := s.accountRepo.Get(c.UserID)
		if err != nil {
			return nil, err
		}

		article, err := s.articleRepo.Get(c.ArticleID)
		if err != nil {
			return nil, err
		}

		comment := GetRes{
			ID:          c.ID,
			Text:        c.Text,
			CreatedDate: c.CreatedDate,
			User:        user,
			Article:     article,
		}

		comments = append(comments, comment)
	}

	return
}

func (s *service) GetArticleComment(ctx context.Context, articleID string) (comments []GetRes, err error) {
	if len(articleID) < 1 {
		err = errInvalidArgument
		return
	}

	allArticleComments, err := s.commentRepo.GetArticleComment(articleID)
	if err != nil {
		return
	}

	for i := range allArticleComments {
		c := &allArticleComments[i]
		user, err := s.accountRepo.Get(c.UserID)
		if err != nil {
			return nil, err
		}

		article, err := s.articleRepo.Get(c.ArticleID)
		if err != nil {
			return nil, err
		}

		comment := GetRes{
			ID:          c.ID,
			Text:        c.Text,
			CreatedDate: c.CreatedDate,
			User:        user,
			Article:     article,
		}

		comments = append(comments, comment)
	}

	return
}

func (s *service) Update(ctx context.Context, id, text string) (err error) {
	if len(id) < 1 || len(text) < 1 {
		err = errInvalidArgument
		return
	}

	return s.commentRepo.Update(id, text)
}

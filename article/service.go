package article

import (
	"context"

	"sample/common/err"
	"sample/model"
)

var (
	errInvalidArgument = err.New(101, "invalid argument")
)

type Service interface {
	GetArticle(ctx context.Context, id string) (article GetRes, err error)
	AddArticle(ctx context.Context, id string, userID string, title string, description string) (err error)
	UpdateArticle(ctx context.Context, id string, title string) (err error)
	GetAllArticle(ctx context.Context) (article []GetRes, err error)
}

type service struct {
	articleRepo model.ArticleRepo
	accountRepo model.AccountRepo
}

func NewService(articleRepo model.ArticleRepo, accountRepo model.AccountRepo) Service {
	return &service{
		articleRepo: articleRepo,
		accountRepo: accountRepo,
	}
}

func (s *service) GetArticle(ctx context.Context, id string) (article GetRes, err error) {
	if len(id) < 1 {
		err = errInvalidArgument
		return
	}

	a, err := s.articleRepo.Get(id)
	if err != nil {
		return
	}

	u, err := s.accountRepo.Get(a.UserID)
	if err != nil {
		return
	}

	article = GetRes{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
		User:        u,
	}

	return
}

func (s *service) AddArticle(ctx context.Context, id string, userID string, title string, description string) (err error) {
	article := model.Article{
		ID:          id,
		Title:       title,
		UserID:      userID,
		Description: description,
	}

	if !article.IsValid() {
		return errInvalidArgument
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

func (s *service) GetAllArticle(ctx context.Context) (article []GetRes, err error) {
	articles, err := s.articleRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for i, a := range articles {

	}
}

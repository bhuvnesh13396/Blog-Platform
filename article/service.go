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
	Get(ctx context.Context, id string) (article GetRes, err error)
	Add(ctx context.Context, id string, userID string, title string, description string) (err error)
	Update(ctx context.Context, id string, title string) (err error)
	List(ctx context.Context) (article []GetRes, err error)
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

func (s *service) Get(ctx context.Context, id string) (article GetRes, err error) {
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

func (s *service) Add(ctx context.Context, id string, userID string, title string, description string) (err error) {
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

func (s *service) Update(ctx context.Context, id string, title string) (err error) {
	if len(id) < 1 || len(title) < 1 {
		err = errInvalidArgument
		return
	}

	return s.articleRepo.Update(id, title)
}

func (s *service) List(ctx context.Context) (article []GetRes, err error) {
	articles, err := s.articleRepo.GetAll()
	if err != nil {
		return
	}

	for i := range articles {
		a := &articles[i]
		user, err := s.accountRepo.Get(a.UserID)
		if err != nil {
			return nil, err
		}

		ar := GetRes{
			ID:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			User:        user,
		}

		article = append(article, ar)
	}

	return
}

package article

import (
	"context"

	"sample/account"
	"sample/common/err"
	"sample/common/id"
	"sample/model"
)

var (
	errInvalidArgument = err.New(101, "invalid argument")
)

type Service interface {
	Get(ctx context.Context, id string) (res GetRes, err error)
	Add(ctx context.Context, userID string, title string, description string) (id string, err error)
	Update(ctx context.Context, id string, title string) (err error)
	List(ctx context.Context) (res []GetRes, err error)
}

type service struct {
	articleRepo model.ArticleRepo
	accountSvc  account.Service
}

func NewService(articleRepo model.ArticleRepo, accountSvc account.Service) Service {
	return &service{
		articleRepo: articleRepo,
		accountSvc:  accountSvc,
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

	u, err := s.accountSvc.Get1(ctx, a.UserID)
	if err != nil {
		return
	}

	article = GetRes{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
		User: User{
			ID:       u.ID,
			Username: u.Username,
			Name:     u.Name,
		},
	}

	return
}

func (s *service) Add(ctx context.Context, userID string, title string, description string) (aid string, err error) {
	article := model.Article{
		ID:          id.New(),
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	if !article.IsValid() {
		return "", errInvalidArgument
	}

	err = s.articleRepo.Add(article)
	if err != nil {
		return
	}

	return article.ID, nil
}

func (s *service) Update(ctx context.Context, id string, title string) (err error) {
	if len(id) < 1 || len(title) < 1 {
		err = errInvalidArgument
		return
	}

	return s.articleRepo.Update(id, title)
}

func (s *service) List(ctx context.Context) (res []GetRes, err error) {
	articles, err := s.articleRepo.GetAll()
	if err != nil {
		return
	}

	for i := range articles {
		a := &articles[i]
		u, err := s.accountSvc.Get1(ctx, a.UserID)
		if err != nil {
			return []GetRes{}, err
		}

		ar := GetRes{
			ID:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			User: User{
				ID:       u.ID,
				Username: u.Username,
				Name:     u.Name,
			},
		}

		res = append(res, ar)
	}

	return
}

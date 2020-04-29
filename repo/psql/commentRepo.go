package psql

import (
	"database/sql"
	"sample/model"
)

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) (*commentRepo, error) {
	query := "CREATE TABLE IF NOT EXISTS comment (id varchar primary key, text varchar, user_id varchar, article_id varchar, created_date varchar)"
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &commentRepo{
		db: db,
	}, nil
}

func (repo *commentRepo) Add(c model.Comment) (err error) {
	query := "INSERT INTO comment (id, text, user_id, article_id, created_date) VALUES ($1, $2, $3, $4, $5)"
	_, err = repo.db.Exec(query, c.ID, c.Text, c.UserID, c.ArticleID, c.CreatedDate)
	return
}

func (repo *commentRepo) Update(id, text string) (err error) {
	query := "UPDATE comment SET text = $2 where ID = $1"
	_, err = repo.db.Exec(query, id, text)
	return
}

func (repo *commentRepo) GetUserComment(userID string) (allUserComments []model.Comment, err error) {
	query := "SELECT * from comment where user_id = $1"
	rows := repo.db.QueryRow(query, userID)
	var c model.Comment
	switch err := rows.Scan(&c.ID, &c.Text, &c.UserID, &c.ArticleID, &c.CreatedDate); err {

	case sql.ErrNoRows:
		return nil, model.ErrCommentNotFound
	case nil:
		allUserComments = append(allUserComments, c)
		// return a, nil
	}

	// var commentsOfUser []model.Comment
	// if err != nil {
	// 	return nil, model.ErrCommentNotFound
	// }

	// defer rows.Close()
	// for rows.Next() {
	// 	var c model.Comment
	// 	err = rows.Scan(&c.ID, &c.Text, &c.UserID, &c.ArticleID, &c.CreatedDate)
	// 	if err != nil {
	// 		fmt.Printf("%+v\v", c)
	// 		return nil, model.ErrCommentNotFound
	// 	}

	// 	commentsOfUser = append(commentsOfUser, c)
	// }

	return
}

func (repo *commentRepo) GetArticleComment(articleID string) (allArticleComments []model.Comment, err error) {
	query := "SELECT * from comment where article_id = $1"
	rows := repo.db.QueryRow(query, articleID)
	var c model.Comment

	switch err := rows.Scan(&c.ID, &c.Text, &c.UserID, &c.ArticleID, &c.CreatedDate); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		allArticleComments = append(allArticleComments, c)
	}

	return

	// rows, err := repo.db.Exec(query, articleID)

	// var commemtsOfArticle []model.Comment

	// if err != nil {
	// 	return nil, model.ErrCommentNotFound
	// }

	// defer rows.Close()
	// for rows.Next() {
	// 	var c model.Comment
	// 	err = rows.Scan(&c.ID, &c.Text, &c.UserID, &c.ArticleID, &c.CreatedDate)
	// 	if err != nil {
	// 		fmt.Printf("%+v\v", a)
	// 		return nil, model.ErrCommentNotFound
	// 	}

	// 	commemtsOfArticle = append(commemtsOfArticle, c)
	// }

	// return commemtsOfArticle, err
}

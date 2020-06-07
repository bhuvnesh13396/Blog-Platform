package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"sample/account"
	"sample/auth"
	"sample/comment"
	"sample/repo/psql"
	"sample/tag"

	_ "github.com/lib/pq"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func errExit(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	connStr := "user=test password=qweqwe dbname=blog sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	accountRepo, err := psql.NewAccountRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	sessionRepo, err := psql.NewSessionRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	articleRepo, err := psql.NewArticleRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	commentRepo, err := psql.NewCommentRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	tagRepo, err := psql.NewTagRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	authService := auth.NewService(sessionRepo, accountRepo)
	authHandler := auth.NewHandler(authService)

	accountService := account.NewService(accountRepo)
	accountService = account.NewAuthService(accountService, authService)
	accountHandler := account.NewHandler(accountService)

	commentService := comment.NewService(commentRepo, articleRepo, accountRepo)
	commentService = comment.NewAuthService(commentService, authService)
	commentHandler := comment.NewHandler(commentService)

	tagService := tag.NewService(tagRepo, articleRepo)
	tagHandler := tag.NewHandler(tagService)

	r := http.NewServeMux()

	r.Handle("/auth", authHandler)
	r.Handle("/auth/", authHandler)
	r.Handle("/account", accountHandler)
	r.Handle("/account/", accountHandler)
	r.Handle("/comment", commentHandler)
	r.Handle("/comment/", commentHandler)
	r.Handle("/tag", tagHandler)
	r.Handle("/tag/", tagHandler)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		errExit(err)
	}
}

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
	"sample/article"
	"sample/repo/psql"

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

	articleRepo, err := psql.NewArticleRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	accountService := account.NewService(accountRepo)
	accountHandler := account.NewHandler(accountService)

	articleService := article.NewService(articleRepo, accountRepo)
	articleHandler := article.NewHandler(articleService)

	r := http.NewServeMux()

	r.Handle("/article", articleHandler)
	r.Handle("/article/", articleHandler)
	r.Handle("/account", accountHandler)
	r.Handle("/account/", accountHandler)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		errExit(err)
	}
}

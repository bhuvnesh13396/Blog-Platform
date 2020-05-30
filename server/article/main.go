package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"sample/account/client"
	"sample/article"
	"sample/auth"
	"sample/common/kit"
	"sample/repo/psql"

	_ "github.com/lib/pq"
)

var (
	accountSvcAddr = "http://localhost:8080"
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

	logger := kit.NewJSONLogger(os.Stdout)

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

	accountSvc, err := client.New(accountSvcAddr, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	authService := auth.NewService(sessionRepo, accountRepo)
	authHandler := auth.NewHandler(authService)

	articleService := article.NewService(articleRepo, accountSvc)
	articleService = article.NewLogService(articleService, kit.LoggerWith(logger, "service", "Article"))
	articleService = article.NewAuthService(articleService, authService)
	articleHandler := article.NewHandler(articleService)

	r := http.NewServeMux()

	r.Handle("/auth", authHandler)
	r.Handle("/auth/", authHandler)
	r.Handle("/article", articleHandler)
	r.Handle("/article/", articleHandler)

	log.Println("listening on", ":8081")
	err = http.ListenAndServe(":8081", r)
	if err != nil {
		errExit(err)
	}
}

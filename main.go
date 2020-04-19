package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	accountPSQLPackage "sample/account/repo/psql"
	accountSerivePackage "sample/account/service"
	articlePSQLPackage "sample/article/repo/psql"
	articleServicePackage "sample/article/service"

	"github.com/gorilla/mux"
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

	accountRepo := accountPSQLPackage.NewAccountRepo(db)
	articleRepo := articlePSQLPackage.NewArticleRepo(db)

	accountService := accountSerivePackage.NewService(accountRepo)
	accountHandler := accountSerivePackage.NewHandler(accountService)

	articleService := articleServicePackage.NewService(articleRepo)
	articleHandler := articleServicePackage.NewHandler(articleService)

	r := mux.NewRouter()
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

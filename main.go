package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"sample/service"
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
	connStr := "user=test password=qweqwe dbname=blog"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	accountRepo := psql.NewAccountRepo(db)

	s := service.NewService(accountRepo)
	handler := service.NewHandler(s)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		errExit(err)
	}
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"sample/kv"
	"sample/repo"
	"sample/service"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func errExit(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	accountDB, err := kv.NewDatabse("account")
	if err != nil {
		errExit(err)
	}
	accountRepo := repo.NewAccountRepo(accountDB)

	s := service.NewService(accountRepo)
	handler := service.NewHandler(s)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		errExit(err)
	}
}

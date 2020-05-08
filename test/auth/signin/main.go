package main

import (
	"context"
	"fmt"
	"net/http"

	"sample/auth/client"
)

var (
	username = "nirav"
	password = "qweqweqwe"
	blogAddr = "http://localhost:8080"
	ctx      = context.Background()
)

func main() {
	authSvc, err := client.New(blogAddr, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	token, err := authSvc.Signin(ctx, username, password)
	if err != nil {
		fmt.Printf("Signin Error: %+v\n", err)
		return
	}

	fmt.Printf("%+v\n", token)
}

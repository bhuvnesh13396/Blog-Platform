package main

import (
	"context"
	"fmt"
	"net/http"

	"sample/article/client"
	"sample/common/auth/token"
)

var (
	authtoken = "FTZDRJdnGzOvbirfJNp2LsVt"
	blogAddr  = "http://localhost:8080"
	ctx       = context.Background()
)

func main() {
	ctx = context.WithValue(ctx, token.ContextKey, authtoken)
	articleSvc, err := client.New(blogAddr, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	_, err = articleSvc.Add(ctx, "qwe 231", "alwkjd", "alkwjdlakw")
	if err != nil {
		fmt.Printf("Add Error: %+v\n", err)
		return
	}

	as, err := articleSvc.List(ctx)
	if err != nil {
		fmt.Printf("Get Error: %+v\n", err)
		return
	}

	for _, a := range as {
		fmt.Printf("%+v\n", a)
	}

}

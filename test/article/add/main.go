package main

import (
	"context"
	"fmt"
	"net/http"

	"sample/article/client"
)

var (
	blogAddr = "http://localhost:8080"
	ctx      = context.Background()
)

func main() {
	articleSvc, err := client.New(blogAddr, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	err = articleSvc.Add(ctx, "qwe", "qwe", "alwkjd", "alkwjdlakw")
	if err != nil {
		fmt.Printf("Add Error: %+v\n", err)
		return
	}

	a, err := articleSvc.Get(ctx, "qwe")
	if err != nil {
		fmt.Printf("Get Error: %+v\n", err)
		return
	}

	fmt.Printf("%+v\n", a)
}

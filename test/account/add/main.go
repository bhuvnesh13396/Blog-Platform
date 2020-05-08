package main

import (
	"context"
	"fmt"
	"net/http"

	"sample/account/client"
)

var (
	blogAddr = "http://localhost:8080"
	ctx      = context.Background()
)

func main() {
	account, err := client.New(blogAddr, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	err = account.Add(ctx, "nirav", "nirav", "qweqweqwe")
	if err != nil {
		fmt.Printf("Add Error: %+v\n", err)
		return
	}
	fmt.Printf("Add complete %e\n", err)


	a, err := account.Get(ctx, "nirav")
	if err != nil {
		fmt.Printf("Get Error: %+v\n", err)
		return
	}

	fmt.Printf("%+v\n", a)
}
